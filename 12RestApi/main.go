package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusBadRequest)
		return
	}
	users = append(users, user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(users)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	if len(users) == 0 {
		http.Error(w, "No users found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(users)

}

func updateUsers(w http.ResponseWriter, r *http.Request) {
	if len(users) == 0 {
		http.Error(w, "No users updated", http.StatusNotFound)
		return
	}
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error updating user", http.StatusBadRequest)
		return
	}
	users[0] = user
	json.NewEncoder(w).Encode(users)
}

func deleteUsers(w http.ResponseWriter, r *http.Request) {
	if len(users) == 0 {
		http.Error(w, "No users deleted", http.StatusNotFound)
		return
	}
	users = users[1:]
	fmt.Fprintln(w, "Deleting users")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
	case http.MethodPost:
		createUser(w, r)
	case http.MethodPut:
		updateUsers(w, r)
	case http.MethodDelete:
		deleteUsers(w, r)
	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/users", userHandler)
	fmt.Println("start server")
	http.ListenAndServe(":8080", nil)
}
