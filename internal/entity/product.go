package entity

import (
	"FULLCYCLE/pkg/entity"
	"errors"
	"time"
)

var (
	ErrIDIsRequired       = errors.New("ID is required")
	ErrInvalidId          = errors.New("invalid Id")
	ErrNameIsRequired     = errors.New("name is required")
	ErrInvalidPrice       = errors.New("invalid price")
	ErrPriceIsRequired    = errors.New("price is required")
	ErrEmailIsRequired    = errors.New("email is required")
	ErrPasswordIsRequired = errors.New("password is required")
)

type Product struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Price    float64   `json:"price"`
	CreateAt time.Time `json:"create_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:       entity.NewID(),
		Name:     name,
		Price:    price,
		CreateAt: time.Now(),
	}
	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil

}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrInvalidId
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
	}
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price <= 0 {
		return ErrInvalidId
	}
	return nil
}
