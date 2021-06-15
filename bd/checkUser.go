package bd

import (
	"context"
	"time"

	"github.com/DavidAbellan/Golang-social-net/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*YaExisteUsuario chequea en la base de datos si existe el mail introducido*/
func YaExisteUsuario(email string) (models.Usuario, bool, string) {
	/*Vamos a crear un contexto para acceder a la base de datos y evitar así cuelgues*/
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	/*hay que hacer un defer para cancelar el timeout que hemos añadido en el context*/
	defer cancel()

	db := MongoCN.Database("gosocialnet")
	col := db.Collection("usuario")

	/*condicion para la consulta*/

	condicion := bson.M{"email": email}
	var resultado models.Usuario

	/*decode almacena en la variable resultado*/
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	/*convertir ObjectId que devuelve Mongo en hexadecimal string*/
	id := resultado.ID.Hex()

	if err != nil {
		return resultado, false, id
	}
	return resultado, true, id

}
