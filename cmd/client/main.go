package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-grud/pkg/api"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := api.NewUserServiceClient(conn)

	var action string

	fmt.Println("Enter one of the actions: create, read:all, read, delete")
	fmt.Scan(&action)

	switch action {
	case "create":
		fmt.Println("Creating schema: firstname lastname email")
		var email string
		var firstName string
		var lastName string

		fmt.Scanf("%s %s %s", &firstName, &lastName,  &email)

		u, err := c.Add(context.Background(), &api.AddRequest{
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
		})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(u)
	case "read:all":
		us, err := c.ReadAll(context.Background(), &api.ReadAllRequest{})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(us)
	case "read":
		fmt.Println("Read schema: id")
		var id int32

		fmt.Scanf("%d", &id)

		u, err := c.Read(context.Background(), &api.ReadRequest{Id:id})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(u)
	case "delete":
		fmt.Println("Read schema: id")
		var id int32

		fmt.Scanf("%d", &id)

		u, err := c.Delete(context.Background(), &api.DeleteRequest{Id:id})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(u)
	default:
		log.Fatal("Invalid action")
	}
}
