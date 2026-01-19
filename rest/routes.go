package rest

import (
	handler "ecommerce/rest/handlers"
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	// Product Routes
	mux.Handle(
		"GET /products", 
		manager.With(
			http.HandlerFunc(handler.GetProducts),
		),
	)
	mux.Handle(
		"POST /products", 
		manager.With(
			http.HandlerFunc(handler.CreateProduct),
			middleware.Auth,
		),
	)
	mux.Handle(
		"GET /products/{id}", 
		manager.With(
			http.HandlerFunc(handler.GetProduct), 
		),
	)
	mux.Handle(
		"PUT /products/{id}", 
		manager.With(
			http.HandlerFunc(handler.UpdateProduct), 
			middleware.Auth,
		),
	)
	mux.Handle(
		"DELETE /products/{id}", 
		manager.With(
			http.HandlerFunc(handler.DeleteProduct), 
			middleware.Auth,
		),
	)

	// User Routes
	mux.Handle(
		"POST /users", 
		manager.With(
			http.HandlerFunc(handler.CreateUser),
		),
	)
	mux.Handle(
		"POST /users/login", 
		manager.With(
			http.HandlerFunc(handler.Login),
		),
	)
}