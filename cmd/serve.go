package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handler"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()
	mux.Handle("GET /products", http.HandlerFunc(handler.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handler.CreateProduct))

	fmt.Println("Server started at localhost:5000")
	globalRouter := global_router.GlobalRouter(mux)
	err := http.ListenAndServe(":5000", globalRouter)
	if err != nil {
		fmt.Println("Failed to start server: ", err)
	}
}