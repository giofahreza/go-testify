package service

import (
	"go-testify/entity"
	"go-testify/repository"
)

type ProductService struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return ProductService{ProductRepository: productRepository}
}

func (s ProductService) GetProducts() ([]entity.Product, error) {
	product, err := s.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s ProductService) GetProductByID(ID int) (entity.Product, error) {
	product, err := s.ProductRepository.FindByID(ID)
	if err != nil {
		return entity.Product{}, err
	}

	return product, nil
}
