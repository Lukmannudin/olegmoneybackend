package main

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/golang-jwt/jwt"
	pb "github.com/olegmoney/proto"
	"github.com/olegmoney/server/config"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *Server) SignUp(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	log.Printf("SignUp was invoked %v\n", in)

	return &pb.UserResponse{
		Name:  "Oleg",
		Email: "oleg@gmail.com",
	}, nil
}

func (s *Server) SignIn(ctx context.Context, payload *pb.UserRequest) (*pb.UserResponse, error) {
	email := payload.Email

	// Get Supabase connection
	supabase := config.GetConnectionSupabse()

	var result []struct {
		Id       int    `json:"id"`
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	// Query the database for the user by email
	errorDb := supabase.DB.From("users").Select("*").Eq("email", email).Execute(&result)

	if errorDb != nil {
		// Handle the error gracefully instead of panicking
		return nil, errorDb
	}

	if len(result) == 0 {
		// Handle the case where no user was found with the given email
		return nil, errors.New("user not found")
	}

	password := result[0].Password

	// Compare the stored password hash with the provided password
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(payload.Password))
	if err != nil {
		// Handle incorrect password
		return nil, errors.New("incorrect password")
	}

	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	token, err := sign.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, errors.New("error when generate token")

	}
	// Password is correct, return a success response
	return &pb.UserResponseLogin{
		Token: token,
	}, nil
}
