package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/maktoobgar/go_template/internal/app"
)

var (
	whichApp = flag.String("app", "both", "defines which app to run")
)

func main() {
	flag.Parse()
	w := strings.ToLower(*whichApp)
	if w == "both" || w == "b" || w == "0" {
		app.Both()
	} else if w == "fiber" || w == "f" || w == "1" {
		app.Fiber()
	} else if w == "grpc" || w == "g" || w == "2" {
		app.Grpc()
	} else {
		fmt.Printf("Invalid flag '%s'\n\n", os.Args[1])
		fmt.Print("Valid values for app:\n\n")
		fmt.Println("* '0', 'b', 'both'  == runs fiber and grpc (default)")
		fmt.Println("* '1', 'f', 'fiber' == runs fiber")
		fmt.Println("* '2', 'g', 'grpc'  == runs grpc")
	}
}
