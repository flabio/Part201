package handlers

import (
	"log"
	"net/http"
	"os"

	"test/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {
	router := mux.NewRouter()
	router.HandleFunc("/listadoCliente/", routers.GetCliente).Methods("GET")
	router.HandleFunc("/compras/{fecha}", routers.GetClienteCompras).Methods("GET")
	router.HandleFunc("/resumen/{fecha}", routers.GetClienteResumen).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8787"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
