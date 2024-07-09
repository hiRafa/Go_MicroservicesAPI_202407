package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	associationspb "associations/associationspb"
)

// AssociationsDBModel implements the UserServiceServer interface
type AssociationsDBModel struct {
	associationspb.UnimplementedAssociationsServiceServer
	db *sql.DB
}

// GetUserById retrieves a user by ID
func (s *AssociationsDBModel) GetAssociationsById(ctx context.Context, req *associationspb.AssociationsByIdRequest) (*associationspb.AssociationsResponse, error) {
	// 's' is our key to the AssociationsDBModel vault
	// 's.db' is accessing a magical database tool from the vault
	// 'QueryRow' is using that tool to search for a specific treasure (user data)
	row := s.db.QueryRow("SELECT id, name, short_name, email FROM associations WHERE id = ?", req.Id)

	// Creating a new treasure chest to hold what we find
	association := &associationspb.AssociationsResponse{} /// the road pointer & is here

	// Filling our new treasure chest with the jewels (data) we found
	err := row.Scan(&association.Id, &association.Name, &association.ShortName, &association.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("association not found")
		}
		return nil, fmt.Errorf("error fetching association: %v", err)
	}

	// Returning the directions (pointer) to our new treasure chest
	return association, nil /// the road pointer was already set previously
	// return &user, nil
}

// CreateUser creates a new user
func (s *AssociationsDBModel) CreateAssociations(ctx context.Context, req *associationspb.CreateAssociationsRequest) (*associationspb.AssociationsResponse, error) {
	query := "INSERT INTO associations (name, short_name, email) VALUES (?, ?, ?)"
	result, err := s.db.Exec(query, req.Name, req.ShortName, req.Email)
	if err != nil {
		log.Printf("Error creating association: %v", err)
		return nil, err
	}

	id, _ := result.LastInsertId()
	association := &associationspb.AssociationsResponse{
		Id:        int32(id),
		Name:      req.Name,
		ShortName: req.ShortName,
		Email:     req.Email,
	}
	return association, nil
}
