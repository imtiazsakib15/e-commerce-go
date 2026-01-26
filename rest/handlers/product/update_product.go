package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ReqUpdateProduct struct {
    ID          int     `json:"id"`
    Title       string  `json:"title"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	productID, err :=strconv.Atoi(r.PathValue("id"))
	if(err != nil){
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var product ReqUpdateProduct
	decode := json.NewDecoder(r.Body)
	error := decode.Decode(&product)
	if error != nil {
		fmt.Println(error)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedProduct, err := h.productRepo.Update(productID, repo.Product{
		Title: product.Title,
		Description: product.Description,
		Price: product.Price,
	})
	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, err)
		return
	}
	if(updatedProduct != nil){
		util.SendData(w, http.StatusOK, updatedProduct)
		return
	}
	util.SendError(w, http.StatusBadRequest, "Product not found")
}
