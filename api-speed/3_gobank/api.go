package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type API struct {
	HOST string
	DB   DB
}

func NewAPI(host string, db DB) *API {

	return &API{
		HOST: host,
		DB:   db,
	}
}

func (a *API) Run() error {
	router := mux.NewRouter()

	router.HandleFunc("/login", makeHTTPHandlerFunc(a.handleLogin))
	router.HandleFunc("/", makeHTTPHandlerFunc(a.home))
	router.HandleFunc("/account", makeHTTPHandlerFunc(a.handleAccount))
	router.HandleFunc("/account/{id}", withJWTAuth(makeHTTPHandlerFunc(a.handleAccountByID), a.DB))
	router.HandleFunc("/transfer", makeHTTPHandlerFunc(a.handleTransfer))

	fmt.Println("api listening: ", a.HOST)
	return http.ListenAndServe(a.HOST, router)
}

func (a *API) home(w http.ResponseWriter, r *http.Request) error {
	token, err := createJWT(&Account{Number: 1})
	if err != nil {
		return err
	}

	return WriteJSON(w, 200, token)
}

// 498081
func (a *API) handleLogin(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return fmt.Errorf("method not allowed %s", r.Method)
	}
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	acc, err := a.DB.GetAccountByNumber(int(req.Number))
	if err != nil {
		return err
	}

	if !acc.ValidPassword(req.Password) {
		return fmt.Errorf("not authenticated")
	}

	token, err := createJWT(acc)
	if err != nil {
		return err
	}

	resp := &LoginResponse{
		Token:  token,
		Number: acc.Number,
	}

	return WriteJSON(w, http.StatusOK, resp)
}

func (a *API) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		a.handleGetAccounts(w, r)
		return nil
	}
	if r.Method == "POST" {
		a.handleCreateAccount(w, r)
		return nil
	}
	return fmt.Errorf("method not allowed")
}

func (a *API) handleAccountByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return a.handleGetAccountByID(w, r)
	}
	if r.Method == "DELETE" {
		return a.handleDeleteAccount(w, r)
	}
	return fmt.Errorf("method is not allowed")
}

func (a *API) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	req := &CreateAccountRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}
	account, err := NewAccount(req.FirstName, req.LastName, req.Password)
	if err != nil {
		return err
	}
	if err := a.DB.CreateAccount(account); err != nil {
		return err
	}

	// token, err := createJWT(account)
	// if err != nil {
	// 	return err
	// }

	return WriteJSON(w, 200, account)
}
func (a *API) handleGetAccounts(w http.ResponseWriter, r *http.Request) error {
	Accounts, err := a.DB.GetAccounts()
	if err != nil {
		WriteJSON(w, 400, err.Error())
		return err
	}
	return WriteJSON(w, 200, Accounts)
}
func (a *API) handleGetAccountByID(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}
	account, err := a.DB.GetAccountByID(id)
	if err != nil {
		return err
	}
	return WriteJSON(w, 200, account)
}
func (a *API) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}
	if err := a.DB.DeleteAccount(id); err != nil {
		return err
	}
	return WriteJSON(w, 200, map[string]int{"deleted": id})
}

func (a *API) handleUpdateAccount(w http.ResponseWriter, r *http.Request) error {
	// not yet
	return nil
}

func (a *API) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	transferReq := &TransferRequest{}
	if err := json.NewDecoder(r.Body).Decode(transferReq); err != nil {
		return err
	}
	defer r.Body.Close()
	return WriteJSON(w, 200, transferReq)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WritePermissionDenied(w http.ResponseWriter) {
	WriteJSON(w, http.StatusForbidden, apiError{Error: "permission denied"})
}

func withJWTAuth(handlerFunc http.HandlerFunc, db DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("x-jwt-token")
		token, err := validateJWT(tokenString)
		if err != nil {
			WritePermissionDenied(w)
			return
		}

		if !token.Valid {
			WritePermissionDenied(w)
			return
		}

		userID, err := getID(r)
		if err != nil {
			WritePermissionDenied(w)
			return
		}
		account, err := db.GetAccountByID(userID)
		if err != nil {
			WritePermissionDenied(w)
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		if account.Number != int64(claims["accountNumber"].(float64)) {
			WritePermissionDenied(w)
			return
		}

		handlerFunc(w, r)
	}
}

func createJWT(account *Account) (string, error) {
	secret := os.Getenv("JWT_SEC")
	// time.Now().Add(1000000000 * time.Second).Unix()
	claims := &jwt.MapClaims{
		"expiresAt":     15000,
		"accountNumber": account.Number,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SEC")

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})

}

type apiFunc func(http.ResponseWriter, *http.Request) error

type apiError struct {
	Error string `json:"error"`
}

func makeHTTPHandlerFunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, &apiError{
				Error: err.Error(),
			})
		}
	}
}

func getID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid id given %s", idStr)
	}
	return id, nil
}
