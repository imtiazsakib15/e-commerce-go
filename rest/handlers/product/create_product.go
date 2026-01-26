package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqCreateProduct struct {
    ID          int     `json:"id"`
    Title       string  `json:"title"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct ReqCreateProduct
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	createdProduct, err := h.productRepo.Create(repo.Product{
		Title: newProduct.Title,
		Description: newProduct.Description,
		Price: newProduct.Price,
	})
	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, err)
		return
	}

	util.SendData(w, http.StatusCreated, createdProduct)
}