package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"

	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	operationspb "operations/operationspb"
)

var DB *sql.DB

func main() {
	// Connect to the database (example using SQLite)
	db, err := sql.Open("sqlite", "api.sql")
	if err != nil {
		panic("Database could not connect: " + err.Error())
	}

	DB = db

	err = createTables()
	if err != nil {
		panic("Database could not connect: " + err.Error())
	}
	fmt.Println("Tables created successfully!")

	// Initialize the gRPC server
	listener, err := net.Listen("tcp", ":50061")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server instance
	server := grpc.NewServer()

	// Initialize an instance of your gRPC service implementation
	OperationsDBModel := &OperationsDBModel{db: db}

	// Register the UserService server with the gRPC server
	operationspb.RegisterOperationsServiceServer(server, OperationsDBModel)

	// Register reflection service on gRPC server
	reflection.Register(server)

	log.Println("gRPC server running on port 50061")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	clientGo()
}

func createTables() error {
	createOperationsTable := `
        CREATE TABLE IF NOT EXISTS operations (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            short_name TEXT NOT NULL,
            email TEXT NOT NULL UNIQUE
        )
    `
	_, err := DB.Exec(createOperationsTable)
	if err != nil {
		return fmt.Errorf("could not create operations table: %v", err)
	}

	fmt.Println("Tables created successfully!")
	return nil
}
