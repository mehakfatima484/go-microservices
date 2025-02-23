package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// NotificationService struct defines the notification service with its methods
type NotificationService struct{}

// SendNotification method for NotificationService responds with a notification message
func (s *NotificationService) SendNotification(req string, res *string) error {
	*res = "Notification sent for: " + req
	return nil
}

func main() {
	// Create a new instance of NotificationService and register it with the RPC server
	notificationService := new(NotificationService)
	rpc.Register(notificationService)
	listener, err := net.Listen("tcp", ":8004") // listen on TCP port 8004
	if err != nil {
		fmt.Println("error listening:", err)
		return
	}
	fmt.Println("Notification Service listening on port 8004")

	// Accept incoming RPC requests
	rpc.Accept(listener)
}
