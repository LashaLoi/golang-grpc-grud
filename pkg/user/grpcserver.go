package user

import (
	"context"
	"grpc-grud/pkg/api"
)

// GRPCServer ...
type GRPCServer struct{}

// Add ...
func (s *GRPCServer) Add(ctx context.Context, rew *api.AddRequest) (*api.UserResponse, error) {
	return &api.UserResponse{ User: &api.User{
		Id: 1,
		FirstName: "aliaksei",
		LastName: "loi",
		Email: "lashaloi1409@gmail.com",
	} }, nil
}

// ReadAll ...
func (s *GRPCServer) ReadAll(ctx context.Context, rew *api.ReadAllRequest) (*api.ReadAllResponse, error) {
	// users := []*api.User{{
	// 	Id: 1,
	// 	FirstName: "aliaksei",
	// 	LastName: "loi",
	// 	Email: "lashaloi1409@gmail.com",
	// }}

	// result := &api.ReadAllResponse{ Users: users }

	return nil, nil
}

// Read ...
func (s *GRPCServer) Read(ctx context.Context, rew *api.ReadRequest) (*api.UserResponse, error) {
	return &api.UserResponse{ User: &api.User{
		Id: 1,
		FirstName: "aliaksei",
		LastName: "loi",
		Email: "lashaloi1409@gmail.com",
	} }, nil
}

// Update ...
func (s *GRPCServer) Update(ctx context.Context, rew *api.UpdateRequest) (*api.UserResponse, error) {
	return &api.UserResponse{ User: &api.User{
		Id: 1,
		FirstName: "aliaksei",
		LastName: "loi",
		Email: "lashaloi1409@gmail.com",
	} }, nil
}

// Delete ...
func (s *GRPCServer) Delete(ctx context.Context, rew *api.DeleteRequest) (*api.UserResponse, error) {
	return &api.UserResponse{ User: &api.User{
		Id: 1,
		FirstName: "aliaksei",
		LastName: "loi",
		Email: "lashaloi1409@gmail.com",
	} }, nil
}