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

func Run(s *grpc.Server, ip string, port string) {
	fmt.Printf("gRPC is up and running on %s:%s\n", ip, port)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
