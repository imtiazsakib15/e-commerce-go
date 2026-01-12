package handler

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {

	productID, err :=strconv.Atoi(r.PathValue("productID"))
	if(err != nil){
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for _, product := range database.Products {
		if product.ID == productID {
			util.SendData(w, product, http.StatusOK)
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}