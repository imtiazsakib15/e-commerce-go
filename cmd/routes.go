package cmd

import (
	"ecommerce/handler"
	"ecommerce/middleware"
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
	mux.Handle("GET /products/{productID}", manager.With(
		http.HandlerFunc(handler.GetProductByID), 
		),
	)
}