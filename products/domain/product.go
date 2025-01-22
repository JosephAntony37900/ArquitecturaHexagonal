package domain

//Aquí van las entidades

type Product struct {
	ID    int
	Name  string
	Price float64
}

type ProductRepository interface {
	Save(product Product) error
	FindAll() ([]Product, error)
}

