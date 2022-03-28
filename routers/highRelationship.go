package routers

import (
	"net/http"

	"github.com/iriartico/twittor/db"
	"github.com/iriartico/twittor/models"
)

func HigRelationship(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "The id parameter is mandatory", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserID = UserID
	t.UserRelationID = ID

	status, err := db.InsertRelation(t)
	if err != nil {
		http.Error(w, "An error occurred when inserting the relationship "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "Failed to insert relationship "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
