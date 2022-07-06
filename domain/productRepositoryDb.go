package domain

import (
	"errors"
	"fmt"
	"github.com/dougmendes/grocerylist-backend/logger"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

type ProductRepositoryDb struct {
	client *sqlx.DB
}

func (p ProductRepositoryDb) GetAll() ([]Product, error) {
	products := make([]Product, 0)
	getAllSql := "select * from products"
	err := p.client.Select(&products, getAllSql)
	if err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("%s", http.StatusInternalServerError))
	}
	return products, nil
}

func (product ProductRepositoryDb) GetProductById(id string) (*Product, error) {
	getByIdSql := "select id, product_name , product_desc , brand_id , price, category_id  from products where id = $1"
	var p Product
	log.Println(product.client.Get(&p, getByIdSql, id))
	err := product.client.Get(&p, getByIdSql, id)
	if err != nil {
		logger.Error("Error while scanning customer " + err.Error())
		return nil, errors.New(fmt.Sprintf("%s", http.StatusInternalServerError))
	}
	return &p, nil
}
func NewProductRepositoryDb(client *sqlx.DB) ProductRepositoryDb {
	return ProductRepositoryDb{client: client}
}
