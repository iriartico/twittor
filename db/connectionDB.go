package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*Mongoose => objeto de conexion a la base de datos*/
var Mongoose = ConnectionDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://sysadmin:sysadmin365@clusterdaemon.zqpwr.mongodb.net/db_twittor?retryWrites=true&w=majority")

/*ConnectionDB => conectar a la Base de datos*/
func ConnectionDB() *mongo.Client {
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
