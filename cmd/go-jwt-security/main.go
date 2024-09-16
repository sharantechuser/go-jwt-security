// main.go
package main

import (
	"flag"
	"fmt"

	"github.com/sharantechuser/go-jwt-security/pkg/config"
	"github.com/sharantechuser/go-jwt-security/pkg/router"
)

func main() {

	var port int

	flag.IntVar(&port, "port", 8000, "Specify the port")
	flag.Parse()

	// Access the remaining command-line arguments
	args := flag.Args()
	fmt.Println("Additional arguments:", args)

	config.New()

	r := router.New()

	r.Run(fmt.Sprintf(":%d", port))
}
