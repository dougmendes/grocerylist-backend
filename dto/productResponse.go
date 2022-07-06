package dto

type ProductResponse struct {
	Id          int     `json:"id"`
	Name        string  `json:"product_name"`
	Description string  `json:"product_desc"`
	BrandId     int     `json:"brand_id"`
	Price       float64 `json:"price"`
	CategoryId  int     `json:"category_id"`
}
