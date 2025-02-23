package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	_ "github.com/lib/pq"
)

type User struct {
	Email    string `json:"email" validate:"required,email"` // User's email, required and must be in email format
	Password string `json:"password" validate:"required"`    // User's password, required for login
}

type AuthService struct {
	db       *sql.DB             //database connection
	validate *validator.Validate // Validator instance for struct validation
}

func (s *AuthService) Login(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, `{"message":"Invalid request payload"}`, http.StatusBadRequest)
		return
	}

	// Validate the struct
	err = s.validate.Struct(user)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"message":"Validation failed: %v"}`, err), http.StatusBadRequest)
		return
	}

	var storedUser User
	err = s.db.QueryRow("SELECT email, password FROM users WHERE email = $1", user.Email).Scan(&storedUser.Email, &storedUser.Password)
	if err != nil || storedUser.Password != user.Password {
		http.Error(w, `{"message":"Invalid credentials"}`, http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Email,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // Set token expiration time
	})

	tokenString, err := token.SignedString([]byte("secret")) // Sign the token with a secret key
	if err != nil {
		http.Error(w, `{"message":"Error generating JWT"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString}) // Return the generated JWT token as JSON
}

func main() {
	connStr := "user=postgres password=root123 host=localhost port=5433 dbname=go_tasks sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	fmt.Println("Successfully connected to the database.")

	authService := &AuthService{
		db:       db,
		validate: validator.New(),
	}

	http.HandleFunc("/login", authService.Login) // Register the "/login" route with the Login handler

	fmt.Println("Authentication Service listening on port 8002")
	log.Fatal(http.ListenAndServe(":8002", nil))
}
