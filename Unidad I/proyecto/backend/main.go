package main

import (
	"log"
	"net/http"
	"proyecto/src/routers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	routers.InitDB()

	r := mux.NewRouter()

	r.HandleFunc("/products", routers.GetProducts).Methods("GET")
	r.HandleFunc("/prod/{l}", routers.FilterProducts).Methods("GET")

	corsObj := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(corsObj)(r)))

}
