package db

import (
	"context"
	"fmt"
	"time"

	"github.com/iriartico/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*ReadAllUsers lee todos los usuarios registrados en el sistema, si se recibe R en quienes
trae solo los que contienen relaciones */
func ReadAllUsers(ID string, page int64, search string, typ string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := Mongoose.Database("db_twittor")
	col := db.Collection("users")

	var results []*models.Usuario

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := col.Find(ctx, query, findOptions) //cursor se unsa cuando manejamos Find y no FindOne porque devuelve un cursor para la paginacion
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var found, include bool

	for cursor.Next(ctx) { // Recorriendo el cursor, usuario a usuario
		var s models.Usuario
		err := cursor.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var r models.Relation
		r.UserID = ID
		r.UserRelationID = s.ID.Hex()

		include = false

		found, err = ConsultRelation(r)
		if typ == "new" && found == false {
			include = true
		}
		if typ == "follow" && found == true {
			include = true
		}
		if r.UserRelationID == ID { //verifico que no sea mi propio seguidor
			include = false
		}

		if include == true { // omito las variables que no quiero blanqueandolas
			s.Email = ""
			s.Password = ""
			s.Biography = ""
			s.WebSite = ""
			s.Location = ""
			s.Banner = ""

			results = append(results, &s)
		}
	}

	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	cursor.Close(ctx)
	return results, true
}
