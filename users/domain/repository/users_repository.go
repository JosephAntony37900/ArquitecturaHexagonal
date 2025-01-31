package repository

import "github.com/JosephAntony37900/ArquitecturaHexagonal/users/domain/entities"

type UserRepository interface {
	Save(user entities.User) error
	FindByID(id int) (*entities.User, error)
	FindAll() ([]entities.User, error)
	Update(user entities.User)error
	Delete(id int)error
}