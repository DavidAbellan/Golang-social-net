package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DavidAbellan/Golang-social-net/bd"
)

func ListaUsuarios(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "La página debe ser mayor que cero", http.StatusBadRequest)
		return
	}
	pag := int64(pagTemp)

	result, status := bd.LeoTodosLosUsuarios(IDUsuario, pag, search, typeUser)
	if !status {
		http.Error(w, "Error al recuperar usuarios", http.StatusBadRequest)
		return

	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}
