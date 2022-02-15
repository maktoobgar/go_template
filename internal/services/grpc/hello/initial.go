package hello

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

type server struct {
	UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &HelloReply{Message: "Hello " + in.GetName()}, nil
}

func New(s *grpc.Server) {
	RegisterGreeterServer(s, &server{})
}
