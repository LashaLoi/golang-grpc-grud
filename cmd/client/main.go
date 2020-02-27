package main

import (
	"fmt"
	"log"

	"grpc-grud/pkg/user"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		errors.Wrap(err, "Unable to connect")
	}

	c := user.NewUserServiceClient(conn)

	var action string

	fmt.Println("Enter one of the actions: create, read:all, read, delete")
	fmt.Scan(&action)

	switch action {
	case "create":
		Create(c)
	case "read:all":
		ReadAll(c)
	case "read":
		Read(c)
	case "delete":
		Delete(c)
	default:
		log.Fatal("Invalid action")
	}
}
