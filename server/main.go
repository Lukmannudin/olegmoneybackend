package main

import (
	pb "github.com/olegmoney/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addr = "localhost:50051"

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
