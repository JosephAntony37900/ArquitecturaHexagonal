package application

import (
	"fmt"
	_"github.com/JosephAntony37900/ArquitecturaHexagonal/users/domain/entities"
	"github.com/JosephAntony37900/ArquitecturaHexagonal/users/domain/repository"
)

type UpdateUser struct {
	repo repository.UserRepository
}

func NewUpdateUser(repo repository.UserRepository) *UpdateUser{
	return &UpdateUser{repo: repo}
}

func (us *UpdateUser) Run(id int, name string, email string, password string)error{
	user, err := us.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("usuario no encontrado: %w", err)
	}

	//actualizo los campos del usuario:
	user.Name= name
	user.Email = email
	user.Password= password

	//guardo los cambios en el repositorio:
	if err := us.repo.Update(*user); err != nil {
		return fmt.Errorf("error actualizando el usuario: %w", err)
	}

	return nil
}