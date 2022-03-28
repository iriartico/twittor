package routers

import (
	"net/http"

	"github.com/iriartico/twittor/db"
)

/*RemoveTweet permite eliminar un tweet determinado*/
func RemoveTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Must send the id parameter", http.StatusBadRequest)
		return
	}

	err := db.DeleteTweet(ID, UserID)
	if err != nil {
		http.Error(w, "An error occurred while trying to delete the tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
