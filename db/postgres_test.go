package db

import (
	"log"
	"testing"
)

func TestConn(t *testing.T) {
	db := Conn()
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
