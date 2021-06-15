package routers

import (
	"encoding/json"
	"net/http"

	"github.com/DavidAbellan/Golang-social-net/bd"
	"github.com/DavidAbellan/Golang-social-net/models"
)

/*Registro crea en la BD un usuario*/
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	/*el body de r es un stream , una vez se registra desaparece*/
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos/n"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El campo Email no puede estar vacío/n"+err.Error(), 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "El password debe contener al menos 6 carácteres/n"+err.Error(), 400)
		return
	}
	_, encontrado, _ := bd.YaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "El usuario ya existe(Email registrado)/n"+err.Error(), 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)

	if err != nil {

		http.Error(w, "Ha habido un error en el registro /n"+err.Error(), 400)
		return

	}
	if status == false {
		http.Error(w, "No se ha podido realizar el registro ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
