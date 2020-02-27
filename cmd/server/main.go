package main

import (
	"log"
	"net"

	"grpc-grud/pkg/store"
	"grpc-grud/pkg/user"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func main() {
	db, err := store.Connect()
	if err != nil {
		errors.Wrap(err, "Unable to connect to db")
	}
	defer db.Close()

	s := grpc.NewServer()
	srv := user.NewGRPCServer(db)

	user.RegisterUserServiceServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		errors.Wrap(err, "Unable to connect to port")
	}

	log.Println("Server started")
	if err := s.Serve(l); err != nil {
		errors.Wrap(err, "Unable to start server")
	}
}
