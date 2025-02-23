package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// DataLayerService struct defines the data layer service with its methods
type DataLayerService struct{}

// FetchData method for DataLayerService responds with a data message
func (s *DataLayerService) FetchData(req string, res *string) error {
	*res = "Data fetched for: " + req
	return nil
}

func main() {
	// Create a new instance of DataLayerService and register it with the RPC server
	dataLayerService := new(DataLayerService)
	rpc.Register(dataLayerService)
	listener, err := net.Listen("tcp", ":8010") // listen on TCP port 8010
	if err != nil {
		fmt.Println("error listening:", err)
		return
	}
	fmt.Println("Data Layer Service listening on port 8010")

	// Accept incoming RPC requests
	rpc.Accept(listener)
}
