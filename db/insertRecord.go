package db

import (
	"context"
	"time"

	"github.com/iriartico/twittor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertRecord es la parada final con la DB para insertar los datos de usuario*/
func InsertRecord(u models.Usuario) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() // anula el contexto da de baja el WithTimeout para evitar que se ocupe espacio en el contexto

	db := Mongoose.Database("db_twittor")
	col := db.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
