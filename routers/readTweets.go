package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/iriartico/twittor/db"
)

/*ReadTweets lee los Tweets*/
func ReadTweets(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Must send the id parameter", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Must send the page parameter", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Must send the page parameter with a value greater than zero", http.StatusBadRequest)
		return
	}

	pag := int64(page)
	res, success := db.ReadTweets(ID, pag)
	if success == false {
		http.Error(w, "Error reading tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
