package main

import (
	"log"
	/*los importamos así para que los pueda encontrar Heroku cuando se despliegue*/
	"github.com/DavidAbellan/Golang-social-net/bd"
	"github.com/DavidAbellan/Golang-social-net/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("No se pudo conectar a la BBDD")
		return
	}
	handlers.Manejadores()
}
