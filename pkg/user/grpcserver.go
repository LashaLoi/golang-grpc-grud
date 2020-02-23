package user

import (
	"context"
	"database/sql"
	"grpc-grud/pkg/api"
	"math/rand"
)

// GRPCServer ...
type GRPCServer struct {
	db *sql.DB
}

// NewGRPCServer ...
func NewGRPCServer(db *sql.DB) *GRPCServer {
	return &GRPCServer{
		db: db,
	}
}

// Add ...
func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.UserResponse, error) {
	var id int32
	var email string
	var firstName string
	var lastName string

	if err := s.db.QueryRow("INSERT INTO users (id, email, firstname, lastname) VALUES ($1, $2, $3, $4) RETURNING id, email, firstname, lastname",
		rand.Intn(1000),
		req.GetEmail(),
		req.GetFirstName(),
		req.GetLastName(),
	).Scan(&id, &email, &firstName, &lastName); err != nil {
		return nil, err
	}

	return &api.UserResponse{User: &api.User{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}}, nil
}

// ReadAll ...
func (s *GRPCServer) ReadAll(ctx context.Context, req *api.ReadAllRequest) (*api.ReadAllResponse, error) {
	users := []*api.User{}

	rows, err := s.db.Query(`SELECT id, email, firstname, lastname FROM users`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id int32
		var email string
		var firstName string
		var lastName string

		if err := rows.Scan(&id, &email, &firstName, &lastName); err != nil {
			return nil, err
		}

		user := &api.User{
			Id:        id,
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
		}

		users = append(users, user)
	}

	return &api.ReadAllResponse{Users: users}, nil
}

// Read ...
func (s *GRPCServer) Read(ctx context.Context, req *api.ReadRequest) (*api.UserResponse, error) {
	var email string
	var firstName string
	var lastName string
	var id int32

	if err := s.db.QueryRow("SELECT * FROM users WHERE id = $1",
		req.GetId(),
	).Scan(&id, &lastName, &email, &firstName); err != nil {
		return nil, err
	}

	return &api.UserResponse{User: &api.User{
		Id:        req.GetId(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}}, nil
}

// Delete ...
func (s *GRPCServer) Delete(ctx context.Context, req *api.DeleteRequest) (*api.UserResponse, error) {
	var email string
	var firstName string
	var lastName string

	if err := s.db.QueryRow("DELETE FROM users WHERE id = $1 RETURNING email, firstname, lastname",
		req.GetId()).Scan(&email, &firstName, &lastName); err != nil {
		return nil, err
	}

	return &api.UserResponse{
		User: &api.User{
			Id:        req.GetId(),
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
		},
	}, nil
}
