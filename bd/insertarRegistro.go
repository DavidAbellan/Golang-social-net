package bd

import (
	"context"
	"time"

	"github.com/DavidAbellan/Golang-social-net/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoRegistro se encarga de insertar un usuario en la base de datos*/
func InsertoRegistro(u models.Usuario) (string, bool, error) {

	/*Vamos a crear un contexto para acceder a la base de datos y evitar así cuelgues*/
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	/*hay que hacer un defer para cancelar el timeout que hemos añadido en el context*/
	defer cancel()

	db := MongoCN.Database("gosocialnet")
	col := db.Collection("usuario")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	/*obtener el id del objeto insertado*/

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil

}
