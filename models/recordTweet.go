package models

import "time"

/*RecordTweet es la estructura de registro de los tweets en la DB*/
type RecordTweet struct {
	UserID  string    `bson:"userId" json:"userId,omitempty"`
	Message string    `bson:"message" json:"message,omitempty"`
	Date    time.Time `bson:"date" json:"date,omitempty"`
}
