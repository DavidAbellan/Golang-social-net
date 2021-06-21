package bd

import (
	"context"
	"time"

	"github.com/DavidAbellan/Golang-social-net/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertarTweet(t models.GrabarTweet) (string, bool, error) {

	/*Vamos a crear un contexto para acceder a la base de datos y evitar así cuelgues*/
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	/*hay que hacer un defer para cancelar el timeout que hemos añadido en el context*/
	defer cancel()

	db := MongoCN.Database("gosocialnet")
	col := db.Collection("tweet")

	registro := bson.M{
		"userId":  t.UserId,
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}
	resultado, err := col.InsertOne(ctx, registro)

	if err != nil {
		return "", false, err
	}
	ObjId, _ := resultado.InsertedID.(primitive.ObjectID)
	return ObjId.String(), true, nil

}
