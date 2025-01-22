package main

import (
	"github.com/JosephAntony37900/ArquitecturaHexagonal/products/application"
	"github.com/JosephAntony37900/ArquitecturaHexagonal/products/infrastructure"
	"github.com/JosephAntony37900/ArquitecturaHexagonal/helpers"
	"log"
)

func main() {
	// Conexión a MySQL
	db, err := helpers.NewMySQLConnection()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Crear repositorio e instancia del caso de uso
	productRepo := infrastructure.NewProductRepoMySQL(db)
	createProduct := application.NewCreateProduct(productRepo)

	// Crear un producto
	if err := createProduct.Run("Producto 1", 100.50); err != nil {
		log.Fatalf("Error creating product: %v", err)
	}
}