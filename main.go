package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router = newRouter() //Retorna el objeto de mux con las rutas especificas
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)

}
