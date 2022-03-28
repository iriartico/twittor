package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/iriartico/twittor/db"
)

/*ReadTweetsFollowers lee unicamente los tweets de nuestros seguidores */
func ReadTweetsFollowers(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Must send a page parameter", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Must send a page parameter as integer greater than zero", http.StatusBadRequest)
		return
	}

	res, success := db.ReadTweetsFollowers(UserID, page)
	if success == false {
		http.Error(w, "Error reading tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
