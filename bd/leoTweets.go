package bd

import (
	"context"
	"log"
	"time"

	/*el paquete options permite filtrar y dar un comportamiento
	a la consulta de la BBDD*/
	"github.com/DavidAbellan/Golang-social-net/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*LeoTweets lee los tweets de un usuario concreto*/
func LeoTweets(ID string, pagina int64) ([]*models.DevolverTweets, bool) {
	/*Vamos a crear un contexto para acceder a la base de datos y evitar así cuelgues*/
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	/*hay que hacer un defer para cancelar el timeout que hemos añadido en el context*/
	defer cancel()

	db := MongoCN.Database("gosocialnet")
	col := db.Collection("tweet")

	var result []*models.DevolverTweets

	condicion := bson.M{
		"userid": ID,
	}

	/*PAGINACION CON EL PAQUETE OPTIONS*/
	opciones := options.Find()
	/*límite de la página*/
	opciones.SetLimit(20)
	/*ordenar por fecha*/
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	/*Salto de página (paginación)*/
	opciones.SetSkip((pagina - 1) * 20)
	/*para que salte de 20 tweets en 20*/
	/*pagina 1 - 1 = 0 x 20 = 0 devuelve los 20 primeros*/
	/*pagina 2 - 1 = 1 x 20 = 20 se salta los 20 primeros*/

	cursor, err := col.Find(ctx, condicion, opciones)

	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}
	/*recorremos el array de tweets pasándole un contexto vacio
	porque no nos interesa el timeout*/
	for cursor.Next(context.TODO()) {
		var tweet models.DevolverTweets
		err := cursor.Decode(&tweet)
		if err != nil {
			return result, false
		}
		result = append(result, &tweet)

	}

	return result, true

}
