package main

import (
	"fmt"
	"net"
	"net/rpc"
	"time"
)

// service1 struct defines the service with its methods
type Service1 struct{}

// Hello method for Service1 responds with a greeting
func (s *Service1) Hello(req string, res *string) error {
	*res = "Hello from microservice 1" //set response/greeting message
	return nil
}

func main() {
	// Create a new instance of Service1 and register it with the RPC server
	service1 := new(Service1)
	rpc.Register(service1)
	listener, err := net.Listen("tcp", ":1234") //listen on TCP port 1234
	if err != nil {
		fmt.Println("Error listening:", err) //print error message if listening fails
		return
	}
	fmt.Println("Microservice 1 listening on port 1234") // Print a message indicating that service is listening

	// Go routine to connect to Microservice 2 in a loop until successful
	go func() {
		for {
			client, err := rpc.Dial("tcp", "localhost:5678") //attempt to connect microservice 2
			if err == nil {
				var response string
				err = client.Call("Service2.Hello", "", &response) //call hello method of microservice 2
				if err == nil {
					fmt.Println(response) //print response from microservice 2
					break
				}
			}
			fmt.Println("Retrying connection to Microservice 2...") //print message to indicate error
			time.Sleep(2 * time.Second)                             // Wait for 2 seconds before retrying
		}
	}()
	// Accept incoming RPC requests
	rpc.Accept(listener)
}
