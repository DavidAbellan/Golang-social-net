package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	/*Hay que importar los paquetes locales*/

	"github.com/DavidAbellan/Golang-social-net/middlew"
	"github.com/DavidAbellan/Golang-social-net/routers"
)

/*Manejadores setea el puerto, el handler y pone a escuchar al servidor*/
func Manejadores() {
	router := mux.NewRouter()

	/*EndPoints*/
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
