package db

import (
	"context"
	"log"
	"time"

	"github.com/iriartico/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*ReadTweets lee los tweets de un Perfil */
func ReadTweets(ID string, page int64) ([]*models.ReturnTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := Mongoose.Database("db_twittor")
	col := db.Collection("tweets")

	var results []*models.ReturnTweets

	condition := bson.M{
		"userId": ID,
	}

	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "date", Value: -1}}) //ordena por fecha en orden descendente
	opciones.SetSkip((page - 1) * 20)                  // donde se lleva a cabo la paginacion

	cursor, err := col.Find(ctx, condition, opciones)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for cursor.Next(context.TODO()) { // context.TODO() crea un contexto vacia sin el retardo de 15s

		var record models.ReturnTweets
		err := cursor.Decode(&record)
		if err != nil {
			return results, false
		}
		results = append(results, &record)
	}
	return results, true
}
