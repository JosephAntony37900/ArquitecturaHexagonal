package application

//Aquí van los casos de uso

import (
	"fmt"
	"github.com/JosephAntony37900/ArquitecturaHexagonal/products/domain"
)

type CreateProduct struct {
	repo domain.ProductRepository
}

func NewCreateProduct(repo domain.ProductRepository) *CreateProduct {
	return &CreateProduct{repo: repo}
}

func (cp *CreateProduct) Run(name string, price float64) error {
	product := domain.Product{Name: name, Price: price}
	if err := cp.repo.Save(product); err !=nil{
		return fmt.Errorf("error guardando el producto: %w", err)// %w es usado para envolver el error original
	}
	fmt.Println("Producto creado")
	return nil
}