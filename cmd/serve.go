package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	manager := middleware.NewManager()
	manager.Use(middleware.Logger, middleware.Cors, middleware.Preflight)

	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	fmt.Println("Server started at localhost:5000")
	err := http.ListenAndServe(":5000", wrappedMux)
	if err != nil {
		fmt.Println("Failed to start server: ", err)
	}
}