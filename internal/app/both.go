package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Both() {
	// Run Grpc
	if !fiber.IsChild() {
		err := runGrpcProcess()
		if err != nil {
			log.Fatal(err)
		}
	}

	// Run Fiber
	Fiber()
}
