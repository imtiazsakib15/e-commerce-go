package rest

import (
	"ecommerce/config"
	middleware "ecommerce/rest/middlewares"
	"fmt"
	"net/http"
)

func Start(cnf config.Config) {
	mux := http.NewServeMux()

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