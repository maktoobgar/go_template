package app

import (
	g "github.com/maktoobgar/go_template/internal/global"
	hello "github.com/maktoobgar/go_template/internal/services/grpc/hello"
	"github.com/maktoobgar/go_template/pkg/grpc"
)

func Grpc() {
	s := grpc.New()

	// Register handlers
	hello.New(s)

	grpc.Run(s, g.CFG.Grpc.IP, g.CFG.Grpc.Port)
}
