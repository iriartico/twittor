package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/iriartico/twittor/db"
	"github.com/iriartico/twittor/models"
)

func RecordTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(r.Body).Decode(&message)

	record := models.RecordTweet{
		UserID:  UserID,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := db.InsertsTweet(record)
	if err != nil {
		http.Error(w, "An error ocurred to insert the record. Try again "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "The tweet could not be inserted", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
