package routers

import (
	"encoding/json"
	"net/http"

	"strconv"

	"github.com/DavidAbellan/Golang-social-net/bd"
)

/*LeerTweets lee los tweets*/
func LeerTweets(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "No se encontró el id", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "No se encontró la página", http.StatusBadRequest)
		return

	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))

	if err != nil {
		http.Error(w, "La página enviada es menor que 0 "+err.Error(), http.StatusBadRequest)
		return

	}
	pag := int64(pagina)

	resp, correcto := bd.LeoTweets(ID, pag)
	if correcto == false {
		http.Error(w, "Error al rescuperar los tweets", http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
