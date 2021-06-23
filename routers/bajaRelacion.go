package routers

import (
	"net/http"

	"github.com/DavidAbellan/Golang-social-net/bd"
	"github.com/DavidAbellan/Golang-social-net/models"
)

/*BajaRelacion elimina una relación entre dos users*/
func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Ha ocurrido un error borrando la relacion", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario /*la variable global que creamos, se activa con el token*/
	t.UsuarioRelacionID = ID

	status, err := bd.BorrarRelacion(t)

	if err != nil || !status {
		http.Error(w, "Ha ocurrido un error borrando la relacion "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
