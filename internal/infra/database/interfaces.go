package database

import (
	"FULLCYCLE/internal/entity"
)

type UserInterface interface {
	CreateUser(user *entity.User) error
	FindByEmail(emailid string) (*entity.User, error)
}
type ProductInterface interface {
	Create(product *entity.Product) error
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindById(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
