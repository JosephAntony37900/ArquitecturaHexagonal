package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/JosephAntony37900/ArquitecturaHexagonal/products/application"
	"log"
)

type CreateProductController struct {
	createProduct *application.CreateProduct
}

func NewCreateProductController(createProduct *application.CreateProduct) *CreateProductController {
	return &CreateProductController{createProduct: createProduct}
}

func (c *CreateProductController) Handle(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to create a product")

	// Estructura para decodificar el JSON de la solicitud
	var request struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}

	// Decodificar el cuerpo de la solicitud
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Printf("Error decoding request body: %v", err)
		w.WriteHeader(http.StatusBadRequest) // Código 400
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}
	log.Printf("Creating product: Name=%s, Price=%f", request.Name, request.Price)

	// Ejecutar el caso de uso para crear el producto
	if err := c.createProduct.Run(request.Name, request.Price); err != nil {
		log.Printf("Error creating product: %v", err)
		w.WriteHeader(http.StatusInternalServerError) // Código 500
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Respuesta de éxito
	log.Println("Product created successfully")
	w.WriteHeader(http.StatusCreated) // Código 201
	json.NewEncoder(w).Encode(map[string]string{"message": "product created successfully"})
}