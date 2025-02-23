package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strings"
)

type User struct {
	Name     string
	Email    string
	Password string
	Role     string
	IsActive bool
}

func main() {
	// Connect to the Public API Service
	client, err := rpc.Dial("tcp", "localhost:8005")
	if err != nil {
		fmt.Println("Error connecting to Public API Service:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	// Read user details for registration
	fmt.Println("Enter user details for registration:")
	fmt.Print("Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	fmt.Print("Role: ")
	role, _ := reader.ReadString('\n')
	role = strings.TrimSpace(role)

	fmt.Print("Is Active (true/false): ")
	isActiveStr, _ := reader.ReadString('\n')
	isActive := strings.TrimSpace(isActiveStr) == "true"

	registerUser := User{
		Name:     name,
		Email:    email,
		Password: password,
		Role:     role,
		IsActive: isActive,
	}

	var registerRes string
	err = client.Call("PublicAPIService.RegisterUser", registerUser, &registerRes)
	if err != nil {
		fmt.Println("Error calling RegisterUser:", err)
		return
	}
	fmt.Println("RegisterUser response:", registerRes)

	// Read user details for login
	fmt.Println("Enter user details for login:")
	fmt.Print("Email: ")
	email, _ = reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Password: ")
	password, _ = reader.ReadString('\n')
	password = strings.TrimSpace(password)

	loginUser := User{
		Email:    email,
		Password: password,
	}

	var loginRes string
	err = client.Call("PublicAPIService.LoginUser", loginUser, &loginRes)
	if err != nil {
		fmt.Println("Error calling LoginUser:", err)
		return
	}
	fmt.Println("LoginUser response (JWT):", loginRes)
}