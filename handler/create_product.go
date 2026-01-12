package handler

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	newProduct.ID = len(database.Products) + 1
	database.Products = append(database.Products, newProduct)

	util.SendData(w, newProduct, http.StatusCreated)
}