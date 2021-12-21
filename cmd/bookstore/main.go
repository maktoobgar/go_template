package main

import (
	"fmt"
	"log"

	"github.com/maktoobgar/bookstore/internal/config"
)

// Load all local and project configurations
func init() {
	var cfg config.Config

	if err := config.ReadProjectConfigs(&cfg); err != nil {
		log.Fatalln(err)
	}

	if err := config.ReadLocalConfigs(&cfg); err != nil {
		log.Fatalln(err)
	}

	config.SetConfig(&cfg)
	fmt.Println(cfg)
}

func main() {

}
