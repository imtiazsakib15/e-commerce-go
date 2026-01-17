package rest

import (
	handler "ecommerce/rest/handlers"
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(handler.GetProducts),
		),
	)
	mux.Handle("POST /products", manager.With(
		http.HandlerFunc(handler.CreateProduct),
		),
	)
	mux.Handle("GET /products/{id}", manager.With(
		http.HandlerFunc(handler.GetProductByID), 
		),
	)
}