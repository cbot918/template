package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Qauth struct {
	Addr string
	DB   *PostgresDB
}

func NewQauth(addr string, db *PostgresDB) *Qauth {
	return &Qauth{
		Addr: addr,
		DB:   db,
	}
}

func (q *Qauth) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(q.handleUser))

	fmt.Println("Qauth listen on: ", q.Addr)
	http.ListenAndServe(q.Addr, router)
}

func (q *Qauth) handleUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		q.handleGetUser(w, r)
	}
	if r.Method == "POST" {
		q.handleCreateUser(w, r)
	}
	if r.Method == "PUT" {
		lg("put")
	}
	if r.Method == "DELETE" {
		lg("delete")
	}
	return nil
}

func (q *Qauth) handleGetUser(w http.ResponseWriter, r *http.Request) error {

	users, err := q.DB.ReadUser()
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, err)
	}

	WriteJSON(w, http.StatusOK, users)

	return nil
}

func (q *Qauth) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	createUserRequest := &CreateUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(&createUserRequest); err != nil {
		lg("1")
		WriteJSON(w, http.StatusBadRequest, err)
		return err
	}
	user := NewUser(createUserRequest.Email, createUserRequest.Password)
	if err := q.DB.CreateUser(user); err != nil {
		lg("2")
		WriteJSON(w, http.StatusBadRequest, err)
		return err
	}
	WriteJSON(w, http.StatusOK, user)
	return nil
}

func (q *Qauth) handleGetUserByID() error { return nil }

func (q *Qauth) handleDeleteUser() error { return nil }

type apiFunc func(http.ResponseWriter, *http.Request) error

type apiError struct {
	Error string
}

func makeHTTPHandleFunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r)
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, apiError{Error: err.Error()})
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
