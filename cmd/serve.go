package cmd

import (
	"ecommerce/config"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()
	cnf := config.GetConfig()

	manager := middleware.NewManager()
	manager.Use(middleware.Logger, middleware.Cors, middleware.Preflight)

	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	port := ":" + fmt.Sprint(cnf.HTTPPort)
	fmt.Println("Server started at localhost" + port)
	err := http.ListenAndServe(port, wrappedMux)
	if err != nil {
		fmt.Println("Failed to start server: ", err)
	}
}