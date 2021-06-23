package routers

import (
	"net/http"

	"github.com/DavidAbellan/Golang-social-net/bd"
	"github.com/DavidAbellan/Golang-social-net/models"
)

/*AltaRelacion realiza el registro de la relación entre dos users*/
func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Ha ocurrido un error creando la relación de users", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario /*la variable global que creamos, se activa con el token*/
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)

	if err != nil || !status {
		http.Error(w, "Ha ocurrido un error dando de alta la relación de users"+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
