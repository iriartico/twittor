package routers

import (
	"encoding/json"
	"net/http"

	"github.com/iriartico/twittor/db"
	"github.com/iriartico/twittor/models"
)

/*ModifyProfile me ayuda con el metodo PUT a la coleccion de usuarios*/
func ModifyProfile(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Incorrect data "+err.Error(), 400)
		return
	}

	var status bool
	status, err = db.ModifyRecord(t, UserID)
	if err != nil {
		http.Error(w, "An error occurred while trying to modify the registry. Try again "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Failed to modify user record", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
