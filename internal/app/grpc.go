package app

import (
	"fmt"

	g "github.com/maktoobgar/go_template/internal/global"
	hello "github.com/maktoobgar/go_template/internal/handlers/grpc/hello"
	"github.com/maktoobgar/go_template/pkg/grpc"
)

func Grpc() {
	// Print Info
	info()

	s := grpc.New()

	// Register handlers
	hello.New(s)

	fmt.Println("\n==Grpc Startup==")
	grpc.Run(s, g.CFG.Grpc.IP, g.CFG.Grpc.Port)
}
