package models

/*Tweet captura del Body el mensaje que nos llega*/
type Tweet struct {
	Message string `bson:"message" json:"message"`
}
