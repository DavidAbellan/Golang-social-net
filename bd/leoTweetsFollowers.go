package bd

import (
	"context"
	"time"

	"github.com/DavidAbellan/Golang-social-net/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*LeoTweetsSeguidores devuelve los tweets de los followers del usuario*/
func LeoTweetsSeguidores(ID string, pagina int) ([]models.DevuelvoTweetsFollowers, bool) {
	/*Vamos a crear un contexto para acceder a la base de datos y evitar así cuelgues*/
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	/*hay que hacer un defer para cancelar el timeout que hemos añadido en el context*/
	defer cancel()

	db := MongoCN.Database("gosocialnet")
	col := db.Collection("relacion")

	skip := (pagina - 1) * 20
	condiciones := make([]bson.M, 0)

	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "tweet",
		},
	})
	condiciones = append(condiciones, bson.M{"$unwind": "&tweet"})
	condiciones = append(condiciones, bson.M{"$sort": bson.M{
		/*-1 para que ordene de más nuevo a más viejo*/
		"fecha": -1,
	}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, condiciones)
	var result []models.DevuelvoTweetsFollowers
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}

	return result, true

}
