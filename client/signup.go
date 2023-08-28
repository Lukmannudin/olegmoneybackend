package main

import (
	"context"
	pb "github.com/olegmoney/proto"
	"log"
)

func doSignUp(c pb.UserServiceClient) {
	log.Println("do SignUp was invoked")

	res, err := c.SignUp(context.Background(), &pb.UserRequest{
		Name:     "lebron",
		Email:    "LebronJames@gmail.com",
		Password: "123123",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Email+" "+res.Name)
}
