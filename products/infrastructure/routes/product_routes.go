package routes

import (
    "github.com/JosephAntony37900/ArquitecturaHexagonal/products/infrastructure/controllers"
    "github.com/gin-gonic/gin"
)

func SetupProductRoutes(r *gin.Engine, createProductController *controllers.CreateProductController, getProductsController *controllers.GetProductsController, updateProductController *controllers.UpdateProductController, deleteProductController *controllers.DeleteProductController) {
    // las rutas
    r.POST("/products", createProductController.Handle)
    r.GET("/products", getProductsController.Handle)
    r.PUT("/products/:id", updateProductController.Handle)
	r.DELETE("/delete/product/:id", deleteProductController.Handle)
}
