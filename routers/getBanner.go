package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/iriartico/twittor/db"
)

/*GetBanner envia el Banner al HTTP*/
func GetBanner(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Must send the id parameter", http.StatusBadRequest)
		return
	}

	proflle, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	openFile, err := os.Open("uploads/banners/" + proflle.Banner)
	if err != nil {
		http.Error(w, "Image not found", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, openFile)
	if err != nil {
		http.Error(w, "Error copying image", http.StatusBadRequest)
		return
	}
}
