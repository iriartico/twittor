package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*DeleteTweet borra un Tweet determinado*/
func DeleteTweet(ID string, UserId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := Mongoose.Database("db_twittor")
	col := db.Collection("tweets")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id":    objID,
		"userId": UserId,
	}

	_, err := col.DeleteOne(ctx, condition)
	return err
}
