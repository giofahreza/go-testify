package repository

import (
	"unit-test/entity"
)

type ProductRepository interface {
	FindAll() ([]entity.Product, error)
	FindByID(ID int) (entity.Product, error)
}

type ProductRepositoryImpl struct {
	products []entity.Product
}
