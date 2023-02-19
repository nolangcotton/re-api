package main

import (
	"log"
	"github.com/nolangcotton/resrvc/cmd/api"
	"github.com/nolangcotton/resrvc/db"
)

func main() {
	log.Println("Starting RESRVC App")
	db.Checks()
	api.Server()
}
