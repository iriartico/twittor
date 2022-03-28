package db

import (
	"context"
	"fmt"
	"time"

	"github.com/iriartico/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ConsultRelation consulta la relacion entre dos usuarios*/
func ConsultRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := Mongoose.Database("db_twittor")
	col := db.Collection("relations")

	condition := bson.M{
		"userId":         t.UserID,
		"userRelationId": t.UserRelationID,
	}

	var result models.Relation
	fmt.Println(result)
	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil
}
