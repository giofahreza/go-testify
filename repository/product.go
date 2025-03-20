package repository

import (
	"go-testify/entity"
)

type ProductRepository interface {
	FindAll() ([]entity.Product, error)
	FindByID(ID int) (entity.Product, error)
}

type productRepository struct {
	products []entity.Product
}
