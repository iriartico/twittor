package db

import (
	"context"
	"time"

	"github.com/iriartico/twittor/models"
)

/*InsertRelation graba la relacion en la DB*/
func InsertRelation(t models.Relation) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := Mongoose.Database("db_twittor")
	col := db.Collection("relations")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
