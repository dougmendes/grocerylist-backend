package app

import (
	"encoding/json"
	"github.com/dougmendes/grocerylist-backend/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type ProductHandler struct {
	service service.ProductService
}

func (p *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := p.service.GetAllProducts()
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, err)
		return
	} else {
		writeResponse(w, http.StatusOK, products)
	}

}

func (p *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)
	productId := vars["product_id"]
	product, err := p.service.GetProductById(productId)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, err)
		return
	} else {
		writeResponse(w, http.StatusOK, product)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}
