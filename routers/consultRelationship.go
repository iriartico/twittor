package routers

import (
	"encoding/json"
	"net/http"

	"github.com/iriartico/twittor/db"
	"github.com/iriartico/twittor/models"
)

/*ConsultRelationship chequea si existe relacion entre 2 usuarios*/
func ConsultRelationship(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relation
	t.UserID = UserID
	t.UserRelationID = ID

	var res models.ResponseConsultRelation

	status, err := db.ConsultRelation(t)
	if err != nil || status == false {
		res.Status = false
	} else {
		res.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
