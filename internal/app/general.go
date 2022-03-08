package app

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/gofiber/fiber/v2"
	_ "github.com/maktoobgar/go_template/internal/app/load"
	g "github.com/maktoobgar/go_template/internal/global"
)

var envPreforkChildKey string = "GRPC_PREFORK_CHILD"
var envPreforkChildVal string = "1"

func info() {
	// Ignore child processes
	if fiber.IsChild() || isChild() {
		return
	}

	fmt.Println("\n==System Info==")
	mainOrTest := "test"
	if !g.CFG.Debug {
		mainOrTest = "main"
	}
	for _, database := range g.CFG.Databases {
		if database.Name == mainOrTest {
			fmt.Printf("Main Database: %v, %v, %v:%v\n", database.Type, database.DBName, database.Host, database.Port)
			if database.Type != g.DB.Dialect() {
				log.Fatal("expected database is not assigned as main database")
			}
			break
		}
	}
	fmt.Printf("Debug: %v\n", g.CFG.Debug)
}

// Runs grpc along side with fiber in another process
func runGrpcProcess() error {
	cmd := exec.Command(os.Args[0], "-app=grpc")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// add GRPC child flag into child proc env
	cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", envPreforkChildKey, envPreforkChildVal))
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start a child prefork process, error: %v", err)
	}

	return nil
}

func isChild() bool {
	return os.Getenv(envPreforkChildKey) == envPreforkChildVal
}
