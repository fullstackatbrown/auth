package authentication

import (
	"context"

	"github.com/fullstackatbrown/auth-infrastructure/pkg/pb"
)

type Server struct {
	pb.UnimplementedAuthenticationServer
}

func (s *Server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{}, nil
}
