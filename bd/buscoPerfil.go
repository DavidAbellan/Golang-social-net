package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/DavidAbellan/Golang-social-net/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BuscoPerfil(ID string) (models.Usuario, error) {
	/*Vamos a crear un contexto para acceder a la base de datos y evitar así cuelgues*/
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	/*hay que hacer un defer para cancelar el timeout que hemos añadido en el context*/
	defer cancel()

	db := MongoCN.Database("gosocialnet")
	col := db.Collection("usuario")

	var perfil models.Usuario
	ObjID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": ObjID,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil)

	/*obtenido el perfil seteamos el password porque no
	nos interesa devolverlo */

	perfil.Password = ""

	if err != nil {
		fmt.Println("Usuario no encontrado" + err.Error())
		return perfil, err
	}
	return perfil, nil

}
