package bd

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword encripta el password a través de bcrypt*/
func EncriptarPassword(pass string) (string, error) {
	/*cost es las veces que lo encrypta, cuanto más mejor*/
	cost := 8
	criptedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(criptedPassword), err

}
