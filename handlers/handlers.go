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
	router.HandleFunc("/ver_perfil", middlew.ChequeoBD(middlew.ValidarToken(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificar_perfil", middlew.ChequeoBD(middlew.ValidarToken(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/leo_tweets", middlew.ChequeoBD(middlew.ValidarToken(routers.LeerTweets))).Methods("GET")
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidarToken(routers.GrabarTweet))).Methods("POST")
	router.HandleFunc("/borrar_tweet", middlew.ChequeoBD(middlew.ValidarToken(routers.EliminarTweet))).Methods("DELETE")
	router.HandleFunc("/subir_avatar", middlew.ChequeoBD(middlew.ValidarToken(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtener_avatar", middlew.ChequeoBD(middlew.ValidarToken(routers.ObtenerAvatar))).Methods("GET")
	router.HandleFunc("/subir_banner", middlew.ChequeoBD(middlew.ValidarToken(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtener_banner", middlew.ChequeoBD(middlew.ValidarToken(routers.ObtenerBanner))).Methods("GET")
	router.HandleFunc("/alta_relacion", middlew.ChequeoBD(middlew.ValidarToken(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/baja_relacion", middlew.ChequeoBD(middlew.ValidarToken(routers.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/consulta_relacion", middlew.ChequeoBD(middlew.ValidarToken(routers.ConsultaRelacion))).Methods("GET")
	router.HandleFunc("/lista_usuarios", middlew.ChequeoBD(middlew.ValidarToken(routers.ListaUsuarios))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
