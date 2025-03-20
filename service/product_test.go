package service

import (
	"go-testify/entity"
	"go-testify/repository"
	"testing"

	"errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepo = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = NewProductService(productRepo)

func TestGetProducts(t *testing.T) {
	product := []entity.Product{
		{ID: 1, Name: "Laptop", Price: 1000},
		{ID: 2, Name: "Mouse", Price: 10},
	}

	productRepo.On("FindAll").Return(product, nil)

	result, err := productService.GetProducts()
	assert.Nil(t, err)
	assert.Equal(t, product, result)
}

// Positive case
func TestProductService_GetProductByIDFound(t *testing.T) {
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

// Negative case
func TestProductService_GetProductByIDNotFound(t *testing.T) {
	product := entity.Product{}

	productRepo.On("FindByID", 1).Return(product, errors.New("Product not found"))

	result, err := productService.GetProductByID(1)
	assert.NotNil(t, err)
	assert.Equal(t, product, result)
}
