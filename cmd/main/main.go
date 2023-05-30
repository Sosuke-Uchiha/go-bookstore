package main

import (
	"log"
	"net/http"

	"github.com/Sosuke-Uchiha/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)
	r.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
