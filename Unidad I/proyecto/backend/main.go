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
	r.HandleFunc("/prod/", routers.FilterProducts).Methods("GET")
	r.HandleFunc("/updateProduct/{id}", routers.UpdateProduct).Methods("PUT")

	r.HandleFunc("/orders", routers.GetOrders).Methods("GET")

	corsH := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	corsM := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	corsO := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(corsH, corsM, corsO)(r)))

}
