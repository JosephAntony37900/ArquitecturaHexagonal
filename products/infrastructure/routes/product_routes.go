package routes

import (
	"net/http"
	"github.com/JosephAntony37900/ArquitecturaHexagonal/products/infrastructure/controllers"
	"github.com/gorilla/mux"
)

func SetupProductRoutes(r *mux.Router, createProductController *controllers.CreateProductController) {
	r.HandleFunc("/products", createProductController.Handle).Methods(http.MethodPost)
}