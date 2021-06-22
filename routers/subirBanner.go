package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/DavidAbellan/Golang-social-net/bd"
	"github.com/DavidAbellan/Golang-social-net/models"
)

/*subirBanner sube el banner del usuario*/
func SubirBanner(w http.ResponseWriter, r *http.Request) {
	fila, manejador, err := r.FormFile("banner")
	var extension = strings.Split(manejador.Filename, ".")[1]
	var archivo string = "uploads/banner/" + IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen !! : "+err.Error(), http.StatusBadRequest)
		return
	}
	_, er := io.Copy(f, fila)
	if er != nil {
		http.Error(w, "Error al copiar la imagen !! : "+err.Error(), http.StatusBadRequest)
		return
	}
	var usuario models.Usuario
	var status bool

	usuario.Banner = IDUsuario + "." + extension

	status, err = bd.ModificoRegistro(usuario, IDUsuario)

	if err != nil || !status {
		http.Error(w, "Error grabando la imagen !! : "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
