package rpc

import (
	"google.golang.org/grpc"
)

// Service is an interface that rpc service must implement
type Service interface {
	Register(*grpc.Server) func()
}
