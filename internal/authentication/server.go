package authentication

import (
	"context"
	"log"

	"github.com/fullstackatbrown/auth-infrastructure/pkg/pb"
)

type Server struct {
	pb.UnimplementedAuthenticationServer
}

func (s *Server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Printf("received Login request for user %v\n", in.Email)
	user := &pb.User{
		Id:          "123",
		DisplayName: "John Doe",
		Email:       "johndoe@gmail.com",
		PhoneNumber: "1234567890",
		PhotoURL:    "https://example.com/johndoe.jpg",
		Pronouns:    "he/him",
		MeetingLink: "https://meet.google.com/abc-123",
	}
	return &pb.LoginResponse{User: user}, nil
}
