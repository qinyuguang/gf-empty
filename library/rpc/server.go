package rpc

import (
	"context"
	"errors"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gogf/gf/frame/g"
	"google.golang.org/grpc"
)

const (
	// DefaultAddr is default address when not specify in configure file.
	DefaultAddr = ":50001"
)

var (
	serverOnce sync.Once
	serverIns  *GRPCServer
)

// GRPCServer wraps the grpc.Server and provides more feature.
type GRPCServer struct {
	server    *grpc.Server
	opts      []grpc.ServerOption
	closeChan chan os.Signal
	services  []Service
}

// Server returns an instance of GRPCServer.
func Server() *GRPCServer {
	serverOnce.Do(func() {
		serverIns = &GRPCServer{
			closeChan: make(chan os.Signal),
		}

		// handle signal
		signal.Notify(
			serverIns.closeChan,
			syscall.SIGHUP,
			syscall.SIGINT,
			syscall.SIGQUIT,
			syscall.SIGABRT,
			syscall.SIGKILL,
			syscall.SIGTERM,
			syscall.SIGUSR1,
		)
	})

	return serverIns
}

// Options adds grpc.ServerOption for initialize.
func (s *GRPCServer) Options(opts ...grpc.ServerOption) *GRPCServer {
	s.opts = append(s.opts, opts...)
	return s
}

// RegisterService registers service into GRPCServer.
func (s *GRPCServer) RegisterService(services ...Service) {
	if len(services) == 0 {
		return
	}

	s.services = append(s.services, services...)
}

// Run starts server listening in blocking way.
func (s *GRPCServer) Run(ctx context.Context, wg *sync.WaitGroup) error {
	if err := s.initServer(); err != nil {
		return err
	}

	go s.serve()

	wg.Add(1)
	go s.gracefulShutdown(ctx, wg)

	// blocking using channel.
	<-s.closeChan

	return nil
}

func (s *GRPCServer) initServer() error {
	s.server = grpc.NewServer(s.opts...)

	if len(s.services) == 0 {
		return errors.New("empty service")
	}

	for i := range s.services {
		s.services[i].Register(s.server)()
	}

	return nil
}

func (s *GRPCServer) serve() error {
	address := g.Cfg().GetString("rpc.address", DefaultAddr)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		g.Log().Fatal(err)
		return err
	}
	g.Log().Infof("rpc server start on %s", address)

	if err := s.server.Serve(listener); err != nil {
		g.Log().Fatal(err)
		return err
	}

	return nil
}

func (s *GRPCServer) gracefulShutdown(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			s.server.GracefulStop()
			g.Log().Info("rpc server shutdown")
			return
		}
	}
}
