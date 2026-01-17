package handler

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	productID, err :=strconv.Atoi(r.PathValue("id"))
	if(err != nil){
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var product database.Product
	decode := json.NewDecoder(r.Body)
	error := decode.Decode(&product)
	if error != nil {
		fmt.Println(error)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedProduct := database.Update(productID, product)
	if(updatedProduct != nil){
		util.SendData(w, updatedProduct, http.StatusOK)
		return
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}
