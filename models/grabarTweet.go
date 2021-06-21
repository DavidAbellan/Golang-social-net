package models

import "time"

type GrabarTweet struct {
	UserId  string    `bson:"userid" json:"userid, omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje, omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha, omitempty"`
}
