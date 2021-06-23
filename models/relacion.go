package models

/*Relacion es el modelo para grabar las relaciones de users en la BBDD*/
type Relacion struct {
	UsuarioID         string `bson:"usuarioid" json:"usuarioId"`
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"usuarioRelacionId"`
}
