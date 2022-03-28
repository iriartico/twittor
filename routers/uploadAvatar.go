package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/iriartico/twittor/db"
	"github.com/iriartico/twittor/models"
)

/*UploadAvatar sube el Avatar al servidor*/
func UploadAvatar(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var doc string = "uploads/avatars/" + UserID + "." + extension

	f, err := os.OpenFile(doc, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error uploading image! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error copying image! "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.Usuario
	var status bool

	user.Avatar = UserID + "." + extension
	status, err = db.ModifyRecord(user, UserID)
	if err != nil || status == false {
		http.Error(w, "Error saving to database! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
