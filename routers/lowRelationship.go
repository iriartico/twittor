package routers

import (
	"net/http"

	"github.com/iriartico/twittor/db"
	"github.com/iriartico/twittor/models"
)

/*LowRelation realiza el borrado de la relacion entre usuarios*/
func LowRelationship(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relation
	t.UserID = UserID
	t.UserRelationID = ID

	status, err := db.DeleteRelation(t)
	if err != nil {
		http.Error(w, "An error occurred when removing the relationship "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "Failed to remove relationship "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
