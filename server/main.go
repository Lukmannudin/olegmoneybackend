package main

import (
	"log"
	"net"

	pb "github.com/olegmoney/proto"
	"google.golang.org/grpc"
)

var addr = "localhost:8080"

type Server struct {
	pb.UserServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
