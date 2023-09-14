package service

import (
	"context"
	"fmt"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
)

func NewGrpcServer(host string, port int) (server *GrpcServer, err error) {
	s := &GrpcServer{
		Host: host,
		Port: port,
	}
	return s, nil
}

// GrpcServer gRPC server
type GrpcServer struct {
	Host string
	Port int
}

type server struct{}

func (s *server) Echo(ctx context.Context, arg *Request) (*Response, error) {
	p, _ := peer.FromContext(ctx)
	response := fmt.Sprintf("received message from %s, content is: %s\n", p.Addr.String(), arg.GetMessage())
	return &Response{Message: response}, nil
}

func (s *GrpcServer) Start() {
	gs := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)
	RegisterEchoServiceServer(gs, &server{})
	reflection.Register(gs)
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Host, s.Port))
	if err != nil {
		panic(err)
	}
	if err := gs.Serve(l); err != nil {
		panic(err)
	}
}
