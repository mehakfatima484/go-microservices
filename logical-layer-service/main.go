package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// LogicalLayerService struct defines the logical layer service with its methods
type LogicalLayerService struct{}

// HandleLogic method for LogicalLayerService responds with a logic message
func (s *LogicalLayerService) HandleLogic(req string, res *string) error {
	*res = "Logic handled for: " + req
	return nil
}

func main() {
	// Create a new instance of LogicalLayerService and register it with the RPC server
	logicalLayerService := new(LogicalLayerService)
	rpc.Register(logicalLayerService)
	listener, err := net.Listen("tcp", ":8009") // listen on TCP port 8009
	if err != nil {
		fmt.Println("error listening:", err)
		return
	}
	fmt.Println("Logical Layer Service listening on port 8009")

	// Accept incoming RPC requests
	rpc.Accept(listener)
}
