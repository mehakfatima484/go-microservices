package main

import (
	"fmt"
	"net"
	"net/rpc"
	"time"
)

// service 2 struct defines the service with its methods
type Service2 struct{}

// Hello method for Service2 responds with a greeting
func (s *Service2) Hello(req string, res *string) error {
	*res = "Hello from microservice 2" //response/greeting message
	return nil                         //return nil if no error
}

func main() {
	//creates new instance and register it to RPC server
	service2 := new(Service2)
	rpc.Register(service2)
	//listen to port 5678
	listener, err := net.Listen("tcp", ":5678")
	if err != nil {
		fmt.Println("Error listening:", err) //error message
		return
	}
	fmt.Println("Microservice 2 listening on port 5678")

	// Go routine try to connect to Microservice 1 in a loop
	go func() {
		for {
			client, err := rpc.Dial("tcp", "localhost:1234") // Attempt to connect to Microservice 1
			if err == nil {
				var response string
				err = client.Call("Service1.Hello", "", &response) //call hello method of microservice 1
				if err == nil {
					fmt.Println(response) //print response
					break
				}
			}
			fmt.Println("Retrying connection to Microservice 1...") //print message indicating retry
			time.Sleep(2 * time.Second)                             // 2 second wait time
		}
	}()
	//accept incoming RPC request
	rpc.Accept(listener)
}
