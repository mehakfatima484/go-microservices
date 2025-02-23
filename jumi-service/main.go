package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// JumiService struct defines the Jumi service with its methods
type JumiService struct{}

// HandleJumi method for JumiService responds with a Jumi message
func (s *JumiService) HandleJumi(req string, res *string) error {
	*res = "Jumi handled for: " + req
	return nil
}

func main() {
	// Create a new instance of JumiService and register it with the RPC server
	jumiService := new(JumiService)
	rpc.Register(jumiService)
	listener, err := net.Listen("tcp", ":8007") // listen on TCP port 8007
	if err != nil {
		fmt.Println("error listening:", err)
		return
	}
	fmt.Println("Jumi Service listening on port 8007")

	// Accept incoming RPC requests
	rpc.Accept(listener)
}
