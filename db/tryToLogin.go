package db

import (
	"github.com/iriartico/twittor/models"
	"golang.org/x/crypto/bcrypt"
)

/*TryToLogin realiza un chequeo de login a la DB*/
func TryToLogin(email string, password string) (models.Usuario, bool) {
	user, found, _ := UserAlreadyExists(email)
	if found == false {
		return user, false
	}

	passwordBytes := []byte(password)   // password encriptada
	passwordDB := []byte(user.Password) // password no encriptada
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}
	return user, true

}
