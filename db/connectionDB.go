package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*Mongoose => objeto de conexion a la base de datos*/
var Mongoose = ConnectionDB()

/*ConnectionDB => conectar a la Base de datos*/
func ConnectionDB() *mongo.Client {
	loadEnv()
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	clientOptions := options.Client().ApplyURI("mongodb+srv://" + user + ":" + pass + "@" + host + ".zqpwr.mongodb.net/db_twittor?retryWrites=true&w=majority")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("successfully connected")
	return client
}

/*CheckConnection => verificar la conexion*/
func CheckConnection() int {
	err := Mongoose.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}
}
