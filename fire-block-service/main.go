package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// FireBlocksService struct defines the fire blocks service with its methods
type FireBlocksService struct{}

// ManageFireBlocks method for FireBlocksService responds with a fire blocks message
func (s *FireBlocksService) ManageFireBlocks(req string, res *string) error {
	*res = "Fire blocks managed for: " + req
	return nil
}

func main() {
	// Create a new instance of FireBlocksService and register it with the RPC server
	fireBlocksService := new(FireBlocksService)
	rpc.Register(fireBlocksService)
	listener, err := net.Listen("tcp", ":8006") // listen on TCP port 8006
	if err != nil {
		fmt.Println("error listening:", err)
		return
	}
	fmt.Println("Fire Blocks Service listening on port 8006")

	// Accept incoming RPC requests
	rpc.Accept(listener)
}
