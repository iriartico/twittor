package db

import (
	"context"
	"time"

	"github.com/iriartico/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ModifyRecord permite modificar el perfil de usuario*/
func ModifyRecord(u models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := Mongoose.Database("db_twittor")
	col := db.Collection("users")

	record := make(map[string]interface{})

	/*Validando cada parametro que se envia para actualizar*/
	if len(u.Name) > 0 {
		record["name"] = u.Name
	}
	if len(u.LastName) > 0 {
		record["lastName"] = u.LastName
	}
	record["dateOfBirth"] = u.DateOfBirth
	if len(u.Avatar) > 0 {
		record["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		record["banner"] = u.Banner
	}
	if len(u.Biography) > 0 {
		record["biography"] = u.Biography
	}
	if len(u.Location) > 0 {
		record["location"] = u.Location
	}
	if len(u.WebSite) > 0 {
		record["webSite"] = u.WebSite
	}

	updateString := bson.M{
		"$set": record,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}} // filtro para mandar a la instruccion de MongoDB que compara el ID

	_, err := col.UpdateOne(ctx, filter, updateString) // instruccion de MongoDB que actualiza el usuario
	if err != nil {
		return false, err
	}
	return true, nil
}
