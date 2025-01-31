package application

import (
	"fmt"
	_ "github.com/JosephAntony37900/ArquitecturaHexagonal/products/domain/entities"
	"github.com/JosephAntony37900/ArquitecturaHexagonal/products/domain/repository"
)

type UpdateProduct struct {
	repo repository.ProductRepository
}

func NewUpdateProduct(repo repository.ProductRepository) *UpdateProduct {
	return &UpdateProduct{repo: repo}
}

func (up *UpdateProduct) Run(id int, name string, price float64) error {
	// Verificar si el producto existe
	product, err := up.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("product not found: %w", err)
	}

	// Actualizar los campos del producto
	product.Name = name
	product.Price = price

	// Guardar los cambios en el repositorio
	if err := up.repo.Update(*product); err != nil {
		return fmt.Errorf("error updating product: %w", err)
	}

	return nil
}