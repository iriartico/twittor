package db

import (
	"context"
	"time"

	"github.com/iriartico/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

func UserAlreadyExists(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := Mongoose.Database("db_twittor")
	col := db.Collection("users")

	condition := bson.M{"email": email}

	var result models.Usuario

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
