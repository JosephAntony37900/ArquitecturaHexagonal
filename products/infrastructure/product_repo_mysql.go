package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/JosephAntony37900/ArquitecturaHexagonal/products/domain"
)

//Aquí va la implementación con MySQL, digamos que los controladores

type ProductRepoMySQL struct {
	db *sql.DB
}

func NewProductRepoMySQL(db *sql.DB) *ProductRepoMySQL {
	return &ProductRepoMySQL{db: db}
}

func (r *ProductRepoMySQL) Save(product domain.Product) error {
	query := "INSERT INTO products (name, price) VALUES (?, ?)"
	_, err := r.db.Exec(query, product.Name, product.Price)
	if err != nil {
		return fmt.Errorf("Error insertando el producto: %w", err)
	}
	return nil
}

func (r *ProductRepoMySQL) FindAll() ([]domain.Product, error) {
	query := "SELECT id, name, price FROM products"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching products: %w", err)
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}