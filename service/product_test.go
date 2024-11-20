package service

import (
	"testing"
	"unit-test/entity"
	"unit-test/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepo = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = NewProductService(productRepo)

func TestProductService_GetProductByID(t *testing.T) {
	product := entity.Product{
		ID:    1,
		Name:  "Laptop",
		Price: 1000,
	}

	productRepo.On("FindByID", 1).Return(product, nil)

	result, err := productService.GetProductByID(1)
	assert.Nil(t, err)
	assert.Equal(t, product, result)
}

func TestProductService_GetProductByIDNotFound(t *testing.T) {
	product := entity.Product{}

	productRepo.On("FindByID", 1).Return(product, nil)

	result, err := productService.GetProductByID(1)
	assert.NotNil(t, err)
	assert.Equal(t, product, result)
}
