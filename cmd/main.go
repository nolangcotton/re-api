package main

import (
	"log"

	"github.com/nolangcotton/resrvc/cmd/api"
)

func main() {
	log.Println("Starting RESRVC App")
	api.Server()
}
