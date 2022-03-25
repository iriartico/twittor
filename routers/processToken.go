package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/iriartico/twittor/db"
	"github.com/iriartico/twittor/models"
)

var Email string
var UserID string

/*ProcessToken procesa el token para extraer sus valores*/
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	myPassword := []byte("MasterEngineerDevOps")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer") // le quito la palabra Bearer ["Bearer", "[token]"]
	if len(splitToken) != 2 {
		return claims, false, "", errors.New("Invalid token format")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return myPassword, nil
	})
	if err == nil {
		_, found, _ := db.UserAlreadyExists(claims.Email)
		if found == true {
			Email = claims.Email
			UserID = claims.ID.Hex()
		}
		return claims, found, UserID, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("Invalid Token")
	}

	return claims, false, string(""), err

}
