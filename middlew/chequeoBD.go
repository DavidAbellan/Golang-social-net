package middlew

import (
	"net/http"

	"github.com/DavidAbellan/Golang-social-net/bd"
)

/*el middleware ha de devolver el mismo tipo de dato que recibe*/
/*ChequeoBD (Middleware) permite conocer el estado de la base de datos*/
func ChequeoBD(siguiente http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Conexión perdida con la base de datos", 500)
			return
		}
		siguiente.ServeHTTP(w, r)
	}

}
