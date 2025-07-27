package main

import (
	"log"
	"net/http"

	"github.com/IliaKoshkin/cloud-app/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HelloMuxHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
