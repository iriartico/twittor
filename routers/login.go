package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/iriartico/twittor/db"
	"github.com/iriartico/twittor/jwt"
	"github.com/iriartico/twittor/models"
)

/*Login es la funcion para los usuarios*/
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Invalid username and/or password"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "The user's email is required", 400)
		return
	}

	doc, exists := db.TryToLogin(t.Email, t.Password)
	if exists == false {
		http.Error(w, "Invalid username and/or password", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(doc)
	if err != nil {
		http.Error(w, "An error occurred while generating the token"+err.Error(), 400)
		return
	}

	res := models.AnswerLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)

	/*Grabando la cookies en el navegador desde el backend*/
	expire := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expire,
	})
}
