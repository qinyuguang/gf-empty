package book

import (
	bookpb "gf-empty/proto/service/book"

	"google.golang.org/grpc"
)

type Service struct{}

func (service *Service) Register(s *grpc.Server) func() {
	return func() {
		bookpb.RegisterBookServer(s, new(Book))
	}
}
