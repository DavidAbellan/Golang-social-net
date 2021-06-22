package routers

import (
	"net/http"

	"github.com/DavidAbellan/Golang-social-net/bd"
)

func EliminarTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}
	err := bd.BorrarTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")

}
