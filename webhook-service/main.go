package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// WebhookService struct defines the webhook service with its methods
type WebhookService struct{}

// TriggerWebhook method for WebhookService responds with a webhook message
func (s *WebhookService) TriggerWebhook(req string, res *string) error {
	*res = "Webhook triggered for: " + req
	return nil
}

func main() {
	// Create a new instance of WebhookService and register it with the RPC server
	webhookService := new(WebhookService)
	rpc.Register(webhookService)
	listener, err := net.Listen("tcp", ":8003") // listen on TCP port 8003
	if err != nil {
		fmt.Println("error listening:", err)
		return
	}
	fmt.Println("Webhook Service listening on port 8003")

	// Accept incoming RPC requests
	rpc.Accept(listener)
}
