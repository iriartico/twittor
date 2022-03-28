package db

import (
	"context"
	"time"

	"github.com/iriartico/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ReadTweetsFollowers lee los tweets de los seguidores*/
func ReadTweetsFollowers(ID string, page int) ([]models.ReturnTweetsFollowers, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := Mongoose.Database("db_twittor")
	col := db.Collection("relations")

	skip := (page - 1) * 20

	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"userId": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{ // lookup une tablas
			"from":         "tweets", // relacionamos la tabla tweets
			"localField":   "userRelationId",
			"foreignField": "userId", // parametro de la tabla tweets
			"as":           "tweets", // alias
		}})
	conditions = append(conditions, bson.M{"$unwind": "$tweets"})        // nos ayuda a que nos pasen todos los documentos nos vengan en su mismo formato, no en un array
	conditions = append(conditions, bson.M{"$sort": bson.M{"date": -1}}) // para $sort: 1 ascendente, -1 descendente
	conditions = append(conditions, bson.M{"$skip": skip})               // primero skip antes de limit
	conditions = append(conditions, bson.M{"$limit": 20})                // numero del salto entre registros

	cursor, err := col.Aggregate(ctx, conditions) // nuevo framework Agregate de Mongo
	var result []models.ReturnTweetsFollowers

	err = cursor.All(ctx, &result) // recorre todo el cursor y decodifica la salida en result, sino retorna error en err
	if err != nil {
		return result, false
	}

	return result, true
}
