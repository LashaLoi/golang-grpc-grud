package main

import (
	"grpc-grud/pkg/api"
	"grpc-grud/pkg/user"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := user.GRPCServer{}

	api.RegisterUserServiceServer(s, &srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server started")
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
