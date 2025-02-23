package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	IsActive bool   `json:"is_active"`
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	body, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error marshalling request body", http.StatusInternalServerError)
		return
	}

	resp, err := http.Post("http://localhost:8001/register", "application/json", bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, "Error calling User Service: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(responseBody)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	body, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error marshalling request body", http.StatusInternalServerError)
		return
	}

	resp, err := http.Post("http://localhost:8002/login", "application/json", bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, "Error calling Authentication Service: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(responseBody)
}

func main() {
	http.HandleFunc("/register", RegisterUser)
	http.HandleFunc("/login", LoginUser)

	fmt.Println("Public API Service listening on port 8005")
	log.Fatal(http.ListenAndServe(":8005", nil))
}
