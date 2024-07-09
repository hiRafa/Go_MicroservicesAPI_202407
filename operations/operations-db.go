package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	operationspb "operations/operationspb"
)

// OperationsDBModel implements the UserServiceServer interface
type OperationsDBModel struct {
	operationspb.UnimplementedOperationsServiceServer
	db *sql.DB
}

// GetUserById retrieves a user by ID
func (s *OperationsDBModel) GetOperationsById(ctx context.Context, req *operationspb.OperationsByIdRequest) (*operationspb.OperationsResponse, error) {
	// 's' is our key to the OperationsDBModel vault
	// 's.db' is accessing a magical database tool from the vault
	// 'QueryRow' is using that tool to search for a specific treasure (user data)
	row := s.db.QueryRow("SELECT id, name, short_name, email FROM operations WHERE id = ?", req.Id)

	// Creating a new treasure chest to hold what we find
	operation := &operationspb.OperationsResponse{} /// the road pointer & is here

	// Filling our new treasure chest with the jewels (data) we found
	err := row.Scan(&operation.Id, &operation.Name, &operation.ShortName, &operation.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("operation not found")
		}
		return nil, fmt.Errorf("error fetching operation: %v", err)
	}

	// Returning the directions (pointer) to our new treasure chest
	return operation, nil /// the road pointer was already set previously
	// return &user, nil
}

// CreateUser creates a new user
func (s *OperationsDBModel) CreateOperations(ctx context.Context, req *operationspb.CreateOperationsRequest) (*operationspb.OperationsResponse, error) {
	query := "INSERT INTO operations (name, short_name, email) VALUES (?, ?, ?)"
	result, err := s.db.Exec(query, req.Name, req.ShortName, req.Email)
	if err != nil {
		log.Printf("Error creating operation: %v", err)
		return nil, err
	}

	id, _ := result.LastInsertId()
	operation := &operationspb.OperationsResponse{
		Id:        int32(id),
		Name:      req.Name,
		ShortName: req.ShortName,
		Email:     req.Email,
	}
	return operation, nil
}
