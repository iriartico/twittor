package routers

import (
	"encoding/json"
	"net/http"

	"github.com/iriartico/twittor/db"
	"github.com/iriartico/twittor/models"
)

/*Register es la funcion para crear en la DB el registro de usuarios*/
func Register(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t) //El objeto Body es un Stream (una vez que se lee se destruye)
	if err != nil {
		http.Error(w, "Error in received data"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "The email is required", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "You must specify a password of at least 6 characters", 400)
		return
	}

	/*Verificar que no se logueen con el mismo email*/
	_, found, _ := db.UserAlreadyExists(t.Email)
	if found == true {
		http.Error(w, "There is already a registered user with this email", 400)
		return
	}

	_, status, err := db.InsertRecord(t)
	if err != nil {
		http.Error(w, "Error registering user"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "The user was not registered", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
