package bd

import (
	"context"
	"time"

	"github.com/DavidAbellan/Golang-social-net/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ModificoRegistro actualiza un perfil de la base de datos*/
func ModificoRegistro(u models.Usuario, ID string) (bool, error) {
	/*Vamos a crear un contexto para acceder a la base de datos y evitar así cuelgues*/
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	/*hay que hacer un defer para cancelar el timeout que hemos añadido en el context*/
	defer cancel()

	db := MongoCN.Database("gosocialnet")
	col := db.Collection("usuario")

	registro := make(map[string]interface{})

	if len(u.Nombre) > 0 {
		registro["nombre"] = u.Nombre
	}
	if len(u.Apellidos) > 0 {
		registro["apellidos"] = u.Apellidos
	}
	if len(u.Biografia) > 0 {
		registro["biografia"] = u.Biografia
	}
	if len(u.Email) > 0 {
		registro["email"] = u.Email
	}
	if len(u.Avatar) > 0 {
		registro["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		registro["banner"] = u.Banner
	}
	if len(u.SitioWeb) > 0 {
		registro["sitioWeb"] = u.SitioWeb
	}
	if len(u.Ubicacion) > 0 {
		registro["ubicacion"] = u.Ubicacion
	}
	registro["fechaNacimiento"] = u.FechaNacimiento

	updateString := bson.M{
		"$set": registro,
	}
	var objId, _ = primitive.ObjectIDFromHex(ID)

	filtro := bson.M{"_id": bson.M{"$eq": objId}}

	_, err := col.UpdateOne(ctx, filtro, updateString)

	if err != nil {
		return false, err
	}
	return true, nil

}
