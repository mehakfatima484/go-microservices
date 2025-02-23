package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

type User struct {
	Name            string `json:"name" validate:"required"` //validate tag define validation rules for each field
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
	Role            string `json:"role" validate:"required"`
	IsActive        bool   `json:"is_active"`
}

type UserService struct {
	db       *sql.DB
	validate *validator.Validate //validator instance for struct validation
}

func (s *UserService) Register(w http.ResponseWriter, r *http.Request) {
	var user User                                // Create a User instance to hold incoming JSON data
	err := json.NewDecoder(r.Body).Decode(&user) //decodes json request body
	if err != nil {                              // if there is an error
		http.Error(w, `{"message":"Invalid request"}`, http.StatusBadRequest) //400 bad request
		return
	}

	// Validate the struct
	err = s.validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) { //iterate through validation errors
			var errMsg string
			switch err.Tag() { //Check which validation tag failed
			case "required":
				errMsg = fmt.Sprintf("%s is required", err.Field())
			case "email":
				errMsg = "Invalid email address" //customized error message for invalid email
			case "eqfield":
				errMsg = "Password and Confirm Password do not match"
			default:
				errMsg = fmt.Sprintf("Validation failed on %s", err.Field())
			}
			http.Error(w, fmt.Sprintf(`{"message":"%s"}`, errMsg), http.StatusBadRequest)
			//	return
		}
	}
	//insert user data into database
	_, err = s.db.Exec("INSERT INTO users (name, email, password, role, is_active) VALUES ($1, $2, $3, $4, $5)", user.Name, user.Email, user.Password, user.Role, user.IsActive)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"message":"Error inserting user: %v"}`, err), http.StatusInternalServerError)
		return
	}

	// Simulate sending a welcome email
	welcomeEmail := fmt.Sprintf("Welcome %s! Thank you for registering.", user.Name)
	timestamp := time.Now()

	//insert welcome email into database
	_, err = s.db.Exec("INSERT INTO emails (recipient, subject, body, timestamp) VALUES ($1, $2, $3, $4)", user.Email, "Welcome to Our Service", welcomeEmail, timestamp)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"message":"Error saving welcome email: %v"}`, err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Registration successful"})
}

func main() {
	connStr := "user=postgres password=root123 host=localhost port=5433 dbname=go_tasks sslmode=disable"
	db, err := sql.Open("postgres", connStr) //open connection to the postgres
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err) //log and exit if fail connection
	}

	fmt.Println("Successfully connected to the database.") //write sucessfully connected message

	userService := &UserService{
		db:       db,
		validate: validator.New(), //initialize validator
	}

	http.HandleFunc("/register", userService.Register) //Register the "/register" route with the Register handler

	fmt.Println("User Service listening on port 8001") //print server is start listening.
	log.Fatal(http.ListenAndServe(":8001", nil))
}
