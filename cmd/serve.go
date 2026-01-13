package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handler"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()
	mux.Handle("GET /products", middleware.Logger(http.HandlerFunc(handler.GetProducts)))
	mux.Handle("POST /products", middleware.Logger(http.HandlerFunc(handler.CreateProduct)))
	mux.Handle("GET /products/{productID}", middleware.Logger(http.HandlerFunc(handler.GetProductByID)))

	fmt.Println("Server started at localhost:5000")
	globalRouter := global_router.GlobalRouter(mux)
	err := http.ListenAndServe(":5000", globalRouter)
	if err != nil {
		fmt.Println("Failed to start server: ", err)
	}
}