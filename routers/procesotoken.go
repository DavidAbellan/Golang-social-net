package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/DavidAbellan/Golang-social-net/bd"
	"github.com/DavidAbellan/Golang-social-net/models"
)

/*Email variable que se utilizará en todos los endpoints*/
var Email string

/*IDUsuario variable que se utilizará en todos los endpoints*/
var IDUsuario string

/*ProcesarToken procesa el token desde el router*/
func ProcesarToken(tk string) (*models.Claim, bool, string, error) {
	clave := []byte("GO_social_net")
	claim := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claim, false, string(""), errors.New("Formato de Token inválido")
	}
	tk = strings.TrimSpace(splitToken[1])
	/*Validación del token*/
	tkn, err := jwt.ParseWithClaims(tk, claim, func(token *jwt.Token) (interface{}, error) {
		return clave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.YaExisteUsuario(claim.Email)
		if encontrado == true {
			Email = claim.Email
			IDUsuario = claim.ID.Hex()

		}
		return claim, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claim, false, string(" "), errors.New("Token inválido")

	}
	return claim, false, string(" "), err

}
