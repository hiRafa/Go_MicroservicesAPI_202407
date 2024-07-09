package main

import (
	"context"
	"log"

	associationspb "associations/associationspb"

	"google.golang.org/grpc"
)

func clientGo() {
	// Dial the gRPC server
	conn, err := grpc.NewClient("localhost:50051")
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client instance
	client := associationspb.NewAssociationsServiceClient(conn)

	// Prepare context
	ctx := context.Background()

	// Create a sample request message
	request := &associationspb.CreateAssociationsRequest{
		// Populate your request message fields here
	}

	// Invoke the gRPC method with context and request
	response, err := client.CreateAssociations(ctx, request)
	if err != nil {
		log.Fatalf("Error calling CreateAssociations: %v", err)
	}

	// Handle response from server
	log.Printf("CreateAssociations response: %v", response)
}
