package models

/*Tweet captura del body el mensaje*/
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
