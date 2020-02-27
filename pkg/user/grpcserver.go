package user

import (
	"context"
	"math/rand"

	"database/sql"
)

// GRPCServer server struct
type GRPCServer struct {
	db *sql.DB
}

// NewGRPCServer func for creating new grpc server
func NewGRPCServer(db *sql.DB) *GRPCServer {
	return &GRPCServer{
		db: db,
	}
}

// Add handler for adding user to db
func (s *GRPCServer) Add(ctx context.Context, req *AddRequest) (*UserResponse, error) {
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
	return &UserResponse{User: &User{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}}, nil
}

// ReadAll handler for read all users from db
func (s *GRPCServer) ReadAll(ctx context.Context, req *ReadAllRequest) (*ReadAllResponse, error) {
	var users []*User

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

		user := &User{
			Id:        id,
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
		}

		users = append(users, user)
	}
	return &ReadAllResponse{Users: users}, nil
}

// Read handler for read user by id from db
func (s *GRPCServer) Read(ctx context.Context, req *ReadRequest) (*UserResponse, error) {
	var email string
	var firstName string
	var lastName string
	var id int32

	if err := s.db.QueryRow("SELECT * FROM users WHERE id = $1",
		req.GetId(),
	).Scan(&id, &lastName, &email, &firstName); err != nil {
		return nil, err
	}
	return &UserResponse{User: &User{
		Id:        req.GetId(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}}, nil
}

// Delete handler for deleting user by id from db
func (s *GRPCServer) Delete(ctx context.Context, req *DeleteRequest) (*UserResponse, error) {
	var email string
	var firstName string
	var lastName string

	if err := s.db.QueryRow("DELETE FROM users WHERE id = $1 RETURNING email, firstname, lastname",
		req.GetId()).Scan(&email, &firstName, &lastName); err != nil {
		return nil, err
	}
	return &UserResponse{
		User: &User{
			Id:        req.GetId(),
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
		},
	}, nil
}
