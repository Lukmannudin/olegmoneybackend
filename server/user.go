package main

import (
	"context"
	pb "github.com/olegmoney/proto"
	"log"
)

func (s *Server) SignUp(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	log.Printf("SignUp was invoked %v\n", in)

	return &pb.UserResponse{
		Name:  "Oleg",
		Email: "oleg@gmail.com",
	}, nil
}
