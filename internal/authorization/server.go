package authorization

import (
	"context"
	"log"

	"github.com/fullstackatbrown/auth-infrastructure/pkg/pb"
)

type Server struct {
	pb.UnimplementedAuthorizationServer
}

func (s *Server) GetRoles(ctx context.Context, in *pb.GetRolesRequest) (*pb.GetRolesResponse, error) {
	log.Printf("received GetRoles request for user %v\n", in.Uid)
	roles := []*pb.Role{{Fields: []string{"fall22", "cs200", "uta"}}}
	return &pb.GetRolesResponse{Roles: roles}, nil
}
