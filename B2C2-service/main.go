package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// B2C2Service struct defines the B2C2 service with its methods
type B2C2Service struct{}

// HandleB2C2 method for B2C2Service responds with a B2C2 message
func (s *B2C2Service) HandleB2C2(req string, res *string) error {
	*res = "B2C2 handled for: " + req
	return nil
}

func main() {
	// Create a new instance of B2C2Service and register it with the RPC server
	b2c2Service := new(B2C2Service)
	rpc.Register(b2c2Service)
	listener, err := net.Listen("tcp", ":8008") // listen on TCP port 8008
	if err != nil {
		fmt.Println("error listening:", err)
		return
	}
	fmt.Println("B2C2 Service listening on port 8008")

	// Accept incoming RPC requests
	rpc.Accept(listener)
}
