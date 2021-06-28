package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/DavidAbellan/Golang-social-net/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*LeoTodosLosUsuarios devuelve todos los usuarios con los que hay relación*/
func LeoTodosLosUsuarios(ID string, pagina int64, busqueda string, tipo string) ([]*models.Usuario, bool) {
	/*Vamos a crear un contexto para acceder a la base de datos y evitar así cuelgues*/
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	/*hay que hacer un defer para cancelar el timeout que hemos añadido en el context*/
	defer cancel()

	db := MongoCN.Database("gosocialnet")
	col := db.Collection("usuario")

	var resultado []*models.Usuario

	/*el paquete options nos permite introducir opciones en la query*/

	findOptions := options.Find()
	/*primero setSkip para estableces que numero, en que posición,
	paginación, de registros queremos*/
	findOptions.SetSkip((pagina - 1) * 20)
	/*luego seteamos el límite en el número de registros*/
	findOptions.SetLimit(20)
	/*Hay que hacerlo en este orden , si no no funciona*/

	query := bson.M{
		"nombre": bson.M{
			"$regex": `(?i)` + busqueda,
		},
	}

	cursor, err := col.Find(ctx, query, findOptions)

	if err != nil {
		fmt.Println(err.Error())
		return resultado, false
	}

	var encontrado, incluir bool

	for cursor.Next(ctx) {
		var u models.Usuario
		err := cursor.Decode(&u)
		if err != nil {
			fmt.Println(err.Error())
			return resultado, false

		}

		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = u.ID.Hex()

		incluir = false

		encontrado, err = ConsultoRelacion(r)
		if tipo == "new" && encontrado == false {
			incluir = true
		}
		if tipo == "follow" && encontrado == true {
			incluir = true
		}
		if r.UsuarioRelacionID == ID {
			incluir = false
		}

		if incluir == true {
			u.Banner = ""
			u.Email = ""
			u.SitioWeb = ""
			u.Ubicacion = ""
			u.Password = ""
			u.Biografia = ""

			resultado = append(resultado, &u)

		}

	}
	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return resultado, false
	}

	cursor.Close(ctx)
	return resultado, true

}
