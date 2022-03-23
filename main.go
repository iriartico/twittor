package main

import (
	"log"

	"github.com/iriartico/twittor/db"
	"github.com/iriartico/twittor/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("sin conexion a la DB")
		return
	}
	handlers.Handlers()
}
