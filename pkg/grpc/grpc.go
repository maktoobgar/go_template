package grpc

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func New() *grpc.Server {
	return grpc.NewServer()
}

func Run(s *grpc.Server, port string) {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
