package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = crearConexion()
var clientOptions = options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.f5hq1.mongodb.net/socialnet?retryWrites=true&w=majority")

/*crearConexion permite conectar con la BBDD MongoDB*/
func crearConexion() *mongo.Client {

	/*TODO para crear un contexto sin ningún tipo de restricción(ej . TIMEOUT)*/
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	/*hacemos un ping a la bbdd para ver si está disponible(opcional)*/
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión exitosa")
	return client

}

/*checkConnection hace ping a la bd para comprobar que esté disponible*/
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
