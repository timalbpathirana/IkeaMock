package routes

import (
	"github.com/gorilla/mux"
	"github.com/tbpathirana/go-ikea/pkg/controllers"
)

var RegisterStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/product", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/product/{ProductId}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/product/{ProductId}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/product/{ProductId}", controllers.DeleteProductById).Methods("DELETE")
}
