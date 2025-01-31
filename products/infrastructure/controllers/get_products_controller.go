package controllers

import (
    "github.com/JosephAntony37900/ArquitecturaHexagonal/products/application"
    "github.com/gin-gonic/gin"
    "log"
)

type GetProductsController struct {
    getProducts *application.GetProducts
}

func NewGetProductsController(getProducts *application.GetProducts) *GetProductsController {
    return &GetProductsController{getProducts: getProducts}
}

func (c *GetProductsController) Handle(ctx *gin.Context) {
    log.Println("Received request to get all products")
    // Obtener los productos
    products, err := c.getProducts.Run()
    if err != nil {
        log.Printf("Error fetching products: %v", err)
        ctx.JSON(500, gin.H{"error": err.Error()})
        return
    }

    // Devolver los productos en formato JSON
    log.Printf("Returning %d products", len(products))
    ctx.JSON(200, products)
}
