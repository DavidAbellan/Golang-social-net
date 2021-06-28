package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DavidAbellan/Golang-social-net/bd"
)

/*LeoTweetsDeLosFollowers lee los tweets de los followers*/
func LeoTweetsDeLosFollowers(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "No existe esta página", http.StatusBadRequest)
		return

	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))

	if err != nil {
		http.Error(w, "La página debe ser mayor que 0", http.StatusBadRequest)
		return
	}

	respuesta, correcto := bd.LeoTweetsSeguidores(IDUsuario, pagina)

	if !correcto {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)

}
