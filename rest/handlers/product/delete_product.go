package product

import (
	"ecommerce/util"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	productID, err :=strconv.Atoi(r.PathValue("id"))
	if(err != nil){
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := h.productRepo.Delete(productID)
	if(err != nil){
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, err)
		return
	}
	if(product != nil){
		util.SendData(w, http.StatusOK, product)
		return
	}
	util.SendError(w, http.StatusBadRequest, "Product not found")
}