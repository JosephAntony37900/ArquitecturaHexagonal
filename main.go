package main

import (
    "log"
    "github.com/JosephAntony37900/ArquitecturaHexagonal/helpers"
    "github.com/JosephAntony37900/ArquitecturaHexagonal/products/application"
    "github.com/JosephAntony37900/ArquitecturaHexagonal/products/infrastructure/controllers"
    "github.com/JosephAntony37900/ArquitecturaHexagonal/products/infrastructure/repository"
    "github.com/JosephAntony37900/ArquitecturaHexagonal/products/infrastructure/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    // Conexión a MySQL
    db, err := helpers.NewMySQLConnection()
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }
    defer db.Close()

    // Repositorio
    productRepo := repository.NewProductRepoMySQL(db)

    // Casos de uso
    createProduct := application.NewCreateProduct(productRepo)
    getProducts := application.NewGetProducts(productRepo)
    updateProduct := application.NewUpdateProduct(productRepo)
	deleteProduct := application.NewDeleteProduct(productRepo)

    // Controladores
    createProductController := controllers.NewCreateProductController(createProduct)
    getProductsController := controllers.NewGetProductsController(getProducts)
    updateProductController := controllers.NewUpdateProductController(updateProduct)
	deleteProductController := controllers.NewDeleteProductController(deleteProduct)

    // Configuración del enrutador de Gin
    r := gin.Default()

    // Configurar rutas
    routes.SetupProductRoutes(r, createProductController, getProductsController, updateProductController, deleteProductController)

    // Iniciar servidor
    log.Println("Server started at :8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}
