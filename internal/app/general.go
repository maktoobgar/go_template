package app

import (
	"fmt"
	"log"

	g "github.com/maktoobgar/go_template/internal/global"
)

func info() {
	fmt.Println("\n==System Info==")
	mainOrTest := "test"
	if !g.CFG.Debug {
		mainOrTest = "main"
	}
	for _, database := range g.CFG.Databases {
		if database.Name == mainOrTest {
			if database.Type == "sqlite3" {
				fmt.Printf("Main Database: %v, %v\n", database.Type, database.DBName)
			} else {
				fmt.Printf("Main Database: %v, %v, %v:%v\n", database.Type, database.DBName, database.Host, database.Port)
			}
			if database.Type != g.DB.Dialect() {
				log.Fatal("expected database is not assigned as main database")
			}
			break
		}
	}
	fmt.Printf("Debug: %v\n", g.CFG.Debug)
	fmt.Printf("Address: http://%s:%s\n", g.CFG.Api.IP, g.CFG.Api.Port)
	fmt.Printf("Allowed Origins: %v\n", g.CFG.AllowOrigins)
}
