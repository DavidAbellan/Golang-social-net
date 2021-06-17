package jwt

import (
	"time"

	"github.com/DavidAbellan/Golang-social-net/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*GenerarToken genera un token encriptado para el usuario*/
func GenerarToken(t models.Usuario) (string, error) {

	clave := []byte("GO_social_net")

	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitio_web":        t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	TokenStr, err := token.SignedString(clave)

	if err != nil {
		return TokenStr, err
	}
	return TokenStr, nil

}
