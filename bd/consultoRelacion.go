package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/DavidAbellan/Golang-social-net/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ConsultoRelacion comprueba si ha relacion entre dos usuarios*/
func ConsultoRelacion(t models.Relacion) (bool, error) {
	/*Vamos a crear un contexto para acceder a la base de datos y evitar así cuelgues*/
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	/*hay que hacer un defer para cancelar el timeout que hemos añadido en el context*/
	defer cancel()

	db := MongoCN.Database("gosocialnet")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var resultado models.Relacion
	fmt.Println(resultado)
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil

}
