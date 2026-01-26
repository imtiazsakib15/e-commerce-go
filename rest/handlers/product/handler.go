package product

import (
	"ecommerce/repo"
	middleware "ecommerce/rest/middlewares"
)

type Handler struct {
	productRepo repo.ProductRepo
	middlewares *middleware.Middlewares
}

func NewHandler(productRepo repo.ProductRepo, middlewares *middleware.Middlewares) *Handler {
	return &Handler{
		productRepo : productRepo,
		middlewares: middlewares,
	}
}