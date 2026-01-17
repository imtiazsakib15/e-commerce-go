package handler

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	productID, err :=strconv.Atoi(r.PathValue("id"))
	if(err != nil){
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product := database.Delete(productID)
	if(product != nil){
		util.SendData(w, product, http.StatusOK)
		return
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}