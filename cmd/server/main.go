package main

import (
	"log"
	"net"

	"github.com/fullstackatbrown/auth-infrastructure/internal/authentication"
	"github.com/fullstackatbrown/auth-infrastructure/internal/authorization"
	"github.com/fullstackatbrown/auth-infrastructure/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	log.Printf("starting server on port %v\n", 8000)
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthenticationServer(s, &authentication.Server{})
	pb.RegisterAuthorizationServer(s, &authorization.Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
