package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/iriartico/twittor/db"
)

/*ListUsers lee la lista de todos los usuarios*/
func ListUsers(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pageTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Must send the page parameter as an integer greater than zero", http.StatusBadRequest)
		return
	}

	pag := int64(pageTemp)

	result, status := db.ReadAllUsers(UserID, pag, search, typeUser)
	if status == false {
		http.Error(w, "Error reading users", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}
