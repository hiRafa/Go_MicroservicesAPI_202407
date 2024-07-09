package main

import (
	"context"
	"log"

	operationspb "operations/operationspb"

	"google.golang.org/grpc"
)

func clientGo() {
	// Dial the gRPC server
	conn, err := grpc.NewClient("localhost:50061")
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client instance
	client := operationspb.NewOperationsServiceClient(conn)

	// Prepare context
	ctx := context.Background()

	// Create a sample request message
	request := &operationspb.CreateOperationsRequest{
		// Populate your request message fields here
	}

	// Invoke the gRPC method with context and request
	response, err := client.CreateOperations(ctx, request)
	if err != nil {
		log.Fatalf("Error calling CreateOperations: %v", err)
	}

	// Handle response from server
	log.Printf("CreateOperations response: %v", response)
}
