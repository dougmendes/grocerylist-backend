package domain

import "github.com/dougmendes/grocerylist-backend/dto"

type Product struct {
	Id          int     `db:"id"`
	Name        string  `db:"product_name"`
	Description string  `db:"product_desc"`
	BrandId     int     `db:"brand_id"`
	Price       float64 `db:"price"`
	CategoryId  int     `db:"category_id"`
}

func (p Product) ToDto() dto.ProductResponse {
	return dto.ProductResponse{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		BrandId:     p.BrandId,
		Price:       p.Price,
		CategoryId:  p.CategoryId,
	}

}

type CustomerRepository interface {
	GetAll() ([]Product, error)
	GetProductById(id string) (*Product, error)
}
