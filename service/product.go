package service

import (
	"unit-test/entity"
	"unit-test/repository"
)

type ProductService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return ProductService{productRepository}
}

func (service *ProductService) GetProductByID(ID int) (entity.Product, error) {
	product, err := service.productRepository.FindByID(ID)
	if err != nil {
		return product, err
	}

	return product, nil
}
