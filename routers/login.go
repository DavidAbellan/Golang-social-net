package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/DavidAbellan/Golang-social-net/bd"
	"github.com/DavidAbellan/Golang-social-net/jwt"
	"github.com/DavidAbellan/Golang-social-net/models"
)

/*las funciones que estan en la carpeta routers no devuelven nada, son endpoints*/

/*Login se loguea en el sistema*/
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario o contraseña inválidas "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El campo Email no puede estar vacío/n", 400)
		return
	}

	documento, existe := bd.IntentoLogin(t.Email, t.Password)

	if !existe {
		http.Error(w, "Usuario o contraseña inválidas ", 400)
	}

	jwtLlave, err := jwt.GenerarToken(documento)
	if err != nil {
		http.Error(w, "No se ha podido generar el token  "+err.Error(), 400)
		return
	}

	res := models.RespuestaLogin{
		Token: jwtLlave,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)

	/*cookies*/
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "NombreDeLaCookie",
		Value:   jwtLlave,
		Expires: expirationTime,
	})

}
