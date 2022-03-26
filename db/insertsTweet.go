package db

import (
	"context"
	"time"

	"github.com/iriartico/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertsTweet graba el Tweet en la DB*/
func InsertsTweet(t models.RecordTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := Mongoose.Database("db_twittor")
	col := db.Collection("tweets")

	record := bson.M{
		"userId":  t.UserID,
		"message": t.Message,
		"date":    t.Date,
	}
	result, err := col.InsertOne(ctx, record)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
