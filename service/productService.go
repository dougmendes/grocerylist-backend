package service

import (
	"github.com/dougmendes/grocerylist-backend/domain"
	"github.com/dougmendes/grocerylist-backend/dto"
	"log"
)

type ProductService interface {
	GetAllProducts() ([]dto.ProductResponse, error)
	GetProductById(id string) (*dto.ProductResponse, error)
}

type DefaultProductService struct {
	repo domain.ProductRepositoryDb
}

func (product DefaultProductService) GetAllProducts() ([]dto.ProductResponse, error) {
	p, err := product.repo.GetAll()
	if err != nil {
		return nil, err
	}
	response := make([]dto.ProductResponse, 0)
	for _, pro := range p {
		response = append(response, pro.ToDto())
	}

	return response, nil
}

func (product DefaultProductService) GetProductById(id string) (*dto.ProductResponse, error) {
	log.Println("id: ", id)
	p, err := product.repo.GetProductById(id)
	if err != nil {
		return nil, err
	}

	response := p.ToDto()

	return &response, nil
}

func NewProductService(repository domain.ProductRepositoryDb) DefaultProductService {
	return DefaultProductService{repo: repository}
}
