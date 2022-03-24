package middlewares

import (
	"net/http"

	"github.com/iriartico/twittor/db"
)

/*CheckDB es el middleware que me permite conocer el estado de la DB*/
func CheckDB(next http.HandlerFunc) http.HandlerFunc { // los middlew devuelven los mismos valores que se les pasa
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "lost connection to database", 500)
			return // se mata la peticion
		}
		next.ServeHTTP(w, r)
	}
}
