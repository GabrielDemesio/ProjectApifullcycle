package database

import (
	"FULLCYCLE/internal/entity"
)

type UserInterface interface {
	CreateUser(user *entity.User) error
	FindByEmail(emailid string) (*entity.User, error)
}
