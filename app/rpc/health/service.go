package health

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type Service struct{}

func (service *Service) Register(s *grpc.Server) func() {
	return func() {
		healthpb.RegisterHealthServer(s, health.NewServer())
	}
}
