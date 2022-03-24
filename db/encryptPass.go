package db

import "golang.org/x/crypto/bcrypt"

/*EncryptPassword cifra la password del usuario*/
func EncryptPassword(pass string) (string, error) {
	cost := 8 // es un algoritmo basado de 2 elevado al costo mas pasadas para encriptar la contrasena =256 pasadas
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}
