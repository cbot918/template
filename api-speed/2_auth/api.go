package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

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

	router.HandleFunc("/", makeHTTPHandlerFunc(a.home))
	router.HandleFunc("/user", makeHTTPHandlerFunc(a.handleUser))
	router.HandleFunc("/user/{id}", makeHTTPHandlerFunc(a.handleUserByID))

	fmt.Println("api listening: ", a.HOST)
	return http.ListenAndServe(a.HOST, router)
}

func (a *API) home(w http.ResponseWriter, r *http.Request) error {
	return WriteJSON(w, 200, "home")
}

func (a *API) handleUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		a.handleGetUsers(w, r)
	}
	if r.Method == "POST" {
		a.handleCreateUser(w, r)
	}
	return fmt.Errorf("method not allowed")
}

func (a *API) handleUserByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return a.handleGetUserByID(w, r)
	}
	if r.Method == "DELETE" {
		return a.handleDeleteUser(w, r)
	}
	return fmt.Errorf("method is not allowed")
}

func (a *API) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	createUserParam := &CreateUserParam{}
	if err := json.NewDecoder(r.Body).Decode(&createUserParam); err != nil {
		WriteJSON(w, 400, err.Error())
		return err
	}
	user, err := NewUser(createUserParam.Email, createUserParam.Password)
	if err != nil {
		WriteJSON(w, 400, err.Error())
		return err
	}
	if err := a.DB.CreateUser(user); err != nil {
		WriteJSON(w, 400, err.Error())
		return err
	}
	WriteJSON(w, 200, user)
	return nil
}
func (a *API) handleGetUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := a.DB.GetUsers()
	if err != nil {
		WriteJSON(w, 400, err.Error())
		return err
	}
	WriteJSON(w, 200, users)
	return nil
}
func (a *API) handleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}
	account, err := a.DB.GetUserByID(id)
	if err != nil {
		return err
	}
	return WriteJSON(w, 200, account)
}
func (a *API) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}
	if err := a.DB.DeleteUser(id); err != nil {
		return err
	}
	return WriteJSON(w, 200, map[string]int{"deleted": id})
}

func (a *API) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {
	// not yet
	return nil
}

func (a *API) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
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
