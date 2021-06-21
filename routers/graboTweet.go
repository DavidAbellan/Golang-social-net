package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/DavidAbellan/Golang-social-net/bd"
	"github.com/DavidAbellan/Golang-social-net/models"
)

func GrabarTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet

	err := json.NewDecoder(r.Body).Decode(&mensaje)

	regis := models.GrabarTweet{
		UserId:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertarTweet(regis)

	if err != nil {
		http.Error(w, "se ha producido un error al insertar el tweet "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha podido insertar el tweet ", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
