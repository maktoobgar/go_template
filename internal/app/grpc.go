package app

import (
	_ "github.com/maktoobgar/go_template/internal/app/load"
	g "github.com/maktoobgar/go_template/internal/global"
	hello "github.com/maktoobgar/go_template/internal/handlers/grpc/hello"
	"github.com/maktoobgar/go_template/pkg/grpc"
)

func Grpc() {
	s := grpc.New()

	// Register handlers
	hello.New(s)

	grpc.Run(s, g.CFG.Grpc.IP, g.CFG.Grpc.Port)
}
