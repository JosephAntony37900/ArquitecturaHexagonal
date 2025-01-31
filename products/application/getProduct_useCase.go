package application

import (
	"github.com/JosephAntony37900/ArquitecturaHexagonal/products/domain/entities"
	"github.com/JosephAntony37900/ArquitecturaHexagonal/products/domain/repository"
)

type GetProducts struct {
	repo repository.ProductRepository
}

func NewGetProducts(repo repository.ProductRepository) *GetProducts {
	return &GetProducts{repo: repo}
}

func (gp *GetProducts) Run() ([]entities.Product, error) {
	products, err := gp.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}