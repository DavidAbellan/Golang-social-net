package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/DavidAbellan/Golang-social-net/bd"
)

func ObtenerAvatar(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Error al obtener el id !! ", http.StatusBadRequest)
		return
	}
	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Usuario no encontrado "+err.Error(), http.StatusBadRequest)
		return
	}

	var abrirArchivo, er = os.Open("uploads/avatar/" + perfil.Avatar)

	if er != nil {
		http.Error(w, "Imagen no encontrada"+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, abrirArchivo)

	if err != nil {
		http.Error(w, "Error con la imagen "+err.Error(), http.StatusBadRequest)
		return
	}

}
