package bd

import (
	"context"
	"time"

	"github.com/DavidAbellan/Golang-social-net/models"
)

/*BorrarRelacion elimina una relacion de la BBDD*/
func BorrarRelacion(t models.Relacion) (bool, error) {
	/*Vamos a crear un contexto para acceder a la base de datos y evitar así cuelgues*/
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	/*hay que hacer un defer para cancelar el timeout que hemos añadido en el context*/
	defer cancel()

	db := MongoCN.Database("gosocialnet")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)

	if err != nil {
		return false, err
	}

	return true, nil
}
