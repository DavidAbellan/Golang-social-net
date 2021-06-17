package bd

import (
	"github.com/DavidAbellan/Golang-social-net/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin chequea el login en la base de datos*/
func IntentoLogin(email string, pass string) (models.Usuario, bool) {
	user, existe, _ := YaExisteUsuario(email)
	if !existe {
		return user, false
	}
	passwordParam := []byte(pass)
	passwordBD := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordParam)
	if err != nil {
		return user, false
	}
	return user, true
}
