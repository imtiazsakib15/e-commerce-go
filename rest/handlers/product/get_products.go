package product

import (
	"ecommerce/util"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.productRepo.List()
	if err != nil {
		util.SendError(w, http.StatusBadRequest, err)
		return
	}
	util.SendData(w,  http.StatusOK, products)
}