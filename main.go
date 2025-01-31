package main

import (
	"log"
	"net/http"
	"github.com/JosephAntony37900/ArquitecturaHexagonal/helpers"
	"github.com/JosephAntony37900/ArquitecturaHexagonal/products/application"
	"github.com/JosephAntony37900/ArquitecturaHexagonal/products/infrastructure/controllers"
	"github.com/JosephAntony37900/ArquitecturaHexagonal/products/infrastructure/repository"
	"github.com/JosephAntony37900/ArquitecturaHexagonal/products/infrastructure/routes"
	"github.com/gorilla/mux"
)

func main() {
	// Conexión a MySQL
	db, err := helpers.NewMySQLConnection()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Repositorio y casos de uso
	productRepo := repository.NewProductRepoMySQL(db)
	createProduct := application.NewCreateProduct(productRepo)

	// Controladores
	createProductController := controllers.NewCreateProductController(createProduct)

	// Configuración del router
	r := mux.NewRouter()
	routes.SetupProductRoutes(r, createProductController)

	// Iniciar servidor
	log.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}