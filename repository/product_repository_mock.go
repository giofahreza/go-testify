package repository

import (
	"go-testify/entity"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (r *ProductRepositoryMock) FindAll() ([]entity.Product, error) {
	args := r.Called()
	return args.Get(0).([]entity.Product), args.Error(1)
}

func (r *ProductRepositoryMock) FindByID(ID int) (entity.Product, error) {
	args := r.Called(ID)
	return args.Get(0).(entity.Product), args.Error(1)
}
