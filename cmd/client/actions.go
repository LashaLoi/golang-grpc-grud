package main

import (
	"context"
	"fmt"

	"grpc-grud/pkg/user"

	"github.com/pkg/errors"
)

// Create user
func Create(c user.UserServiceClient) {
	fmt.Println("Creating schema: firstname lastname email")
	var email string
	var firstName string

	var lastName string

	fmt.Scanf("%s %s %s", &firstName, &lastName, &email)

	u, err := c.Add(context.Background(), &user.AddRequest{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	})
	if err != nil {
		errors.Wrap(err, "Unable to create a user")
	}

	fmt.Println(u)
}

// ReadAll users
func ReadAll(c user.UserServiceClient) {
	us, err := c.ReadAll(context.Background(), &user.ReadAllRequest{})
	if err != nil {
		errors.Wrap(err, "Unable to read users")
	}

	fmt.Println(us)
}

// Read user by id
func Read(c user.UserServiceClient) {
	fmt.Println("Read schema: id")
	var id int32

	fmt.Scanf("%d", &id)

	u, err := c.Read(context.Background(), &user.ReadRequest{Id: id})
	if err != nil {
		errors.Wrap(err, "Unable to read user by id")
	}

	fmt.Println(u)
}

// Delete user by id
func Delete(c user.UserServiceClient) {
	fmt.Println("Read schema: id")
	var id int32

	fmt.Scanf("%d", &id)

	u, err := c.Delete(context.Background(), &user.DeleteRequest{Id: id})
	if err != nil {
		errors.Wrap(err, "Unable to delete user by id")
	}

	fmt.Println(u)
}
