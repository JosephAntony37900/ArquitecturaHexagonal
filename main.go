package main

import (
    "log"
    "github.com/JosephAntony37900/ArquitecturaHexagonal/helpers"
    productApp "github.com/JosephAntony37900/ArquitecturaHexagonal/products/application"
    productController "github.com/JosephAntony37900/ArquitecturaHexagonal/products/infrastructure/controllers"
    productRepo "github.com/JosephAntony37900/ArquitecturaHexagonal/products/infrastructure/repository"
    productRoutes "github.com/JosephAntony37900/ArquitecturaHexagonal/products/infrastructure/routes"
    userApp "github.com/JosephAntony37900/ArquitecturaHexagonal/users/application"
    userController "github.com/JosephAntony37900/ArquitecturaHexagonal/users/infrastructure/controllers"
    userRepo "github.com/JosephAntony37900/ArquitecturaHexagonal/users/infrastructure/repository"
    userRoutes "github.com/JosephAntony37900/ArquitecturaHexagonal/users/infrastructure/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    // Conexión a MySQL
    db, err := helpers.NewMySQLConnection()
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }
    defer db.Close()

    // Repositorios
    productRepository := productRepo.NewProductRepoMySQL(db)
    userRepository := userRepo.NewCreateUserRepoMySQL(db)

    // Casos de uso para productos
    createProduct := productApp.NewCreateProduct(productRepository)
    getProducts := productApp.NewGetProducts(productRepository)
    updateProduct := productApp.NewUpdateProduct(productRepository)
    deleteProduct := productApp.NewDeleteProduct(productRepository)

    // Casos de uso para usuarios
    createUser := userApp.NewCreateUser(userRepository)

    // Controladores para productos
    createProductController := productController.NewCreateProductController(createProduct)
    getProductsController := productController.NewGetProductsController(getProducts)
    updateProductController := productController.NewUpdateProductController(updateProduct)
    deleteProductController := productController.NewDeleteProductController(deleteProduct)

    // Controladores para usuarios
    createUserController := userController.NewCreateUserController(createUser)

    // Configuración del enrutador de Gin
    r := gin.Default()

    // Configurar rutas de productos
    productRoutes.SetupProductRoutes(r, createProductController, getProductsController, updateProductController, deleteProductController)

    // Configurar rutas de usuarios
    userRoutes.SetupUserRoutes(r, createUserController)

    // Iniciar servidor
    log.Println("Server started at :8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}
