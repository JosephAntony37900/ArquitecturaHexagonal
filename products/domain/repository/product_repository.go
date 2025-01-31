package repository

import "github.com/JosephAntony37900/ArquitecturaHexagonal/products/domain/entities"

type ProductRepository interface {
	Save(product entities.Product) error
	FindByID(id int) (*entities.Product, error)
	FindAll() ([]entities.Product, error)
	Update(product entities.Product) error
	Delete(id int) error
}