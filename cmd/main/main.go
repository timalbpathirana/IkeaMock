package main

import (
	"github.com/gorilla/mux"
	"github.com/tbpathirana/go-ikea/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	// Giving authority to ikea-routers to handle the routing
	routes.RegisterStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:80", r))
}
