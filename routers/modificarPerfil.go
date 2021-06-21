package routers

import (
	"encoding/json"
	"net/http"

	"github.com/DavidAbellan/Golang-social-net/bd"
	"github.com/DavidAbellan/Golang-social-net/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Datos Incorrectos "+err.Error(), 400)
		return
	}
	status, er := bd.ModificoRegistro(t, IDUsuario)

	if er != nil {
		http.Error(w, "Ha habido un error actualizando el perfil "+er.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado encontrar el usuario "+err.Error(), 400)
		return

	}
	w.WriteHeader(http.StatusCreated)
}
