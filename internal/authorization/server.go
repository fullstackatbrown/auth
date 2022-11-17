package authorization

import (
	"context"

	"github.com/fullstackatbrown/auth-infrastructure/pkg/pb"
)

type Server struct {
	pb.UnimplementedAuthorizationServer
}

func (s *Server) GetRoles(ctx context.Context, in *pb.GetRolesRequest) (*pb.GetRolesResponse, error) {
	return nil, nil
}
