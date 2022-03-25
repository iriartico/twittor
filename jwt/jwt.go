package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/iriartico/twittor/models"
)

func GenerateJWT(t models.Usuario) (string, error) {

	myPassword := []byte("MasterEngineerDevOps")

	payload := jwt.MapClaims{
		"email":       t.Email,
		"name":        t.Name,
		"lastName":    t.LastName,
		"dateOfBirth": t.DateOfBirth,
		"biography":   t.Biography,
		"location":    t.Location,
		"webSite":     t.WebSite,
		"_id":         t.ID.Hex(),
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myPassword)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil

}
