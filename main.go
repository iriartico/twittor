package main

import (
	"log"

	"github.com/iriartico/twittor/db"
	"github.com/iriartico/twittor/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("no connection to the DB")
		return
	}
	handlers.Handlers()
}
