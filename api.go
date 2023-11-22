package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
  storage    Storage
}

func NewAPIServer(listenAddr string, storage Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
    storage:    storage,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
  router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccountById))

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {

	switch r.Method {
	case http.MethodGet:
		return s.handleGetAccount(w, r)
	case http.MethodPost:
		return s.handleCreateAccount(w, r)
	case http.MethodPut:
		return s.handleUpdateAccount(w, r)
	case http.MethodDelete:
		return s.handleDeleteAccount(w, r)
	default:
		return fmt.Errorf("Method not allowed %s", r.Method)
	}
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
  accounts, err := s.storage.GetAccounts()
  if err != nil {
    return nil
  }

  return WriteJSON(w, http.StatusOK, accounts)
}

func (s *APIServer) handleGetAccountById( w http.ResponseWriter, r *http.Request) error {
  id := mux.Vars(r)["id"]

  fmt.Println(id)

  return WriteJSON(w, http.StatusOK, &Account{})
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
  createAccountReq := new(CreateAccountRequest)
  if err := json.NewDecoder(r.Body).Decode(createAccountReq); err != nil {
    return err
  }

  account := NewAccount(createAccountReq.FirstName, createAccountReq.LastName)
  if err := s.storage.CreateAccount(account); err != nil {
    return err
  }
  return WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleUpdateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
  w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

