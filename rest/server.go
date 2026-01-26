package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
	"fmt"
	"net/http"
)

type Server struct {
	cnf *config.Config
	userHandler *user.Handler
	productHandler *product.Handler
}

func NewServer(cnf *config.Config, userHandler *user.Handler, productHandler *product.Handler) *Server {
	return &Server{
		cnf: cnf,
		userHandler: userHandler,
		productHandler: productHandler,
	}
}

func (server *Server) Start() {
	mux := http.NewServeMux()

	manager := middleware.NewManager()
	manager.Use(middleware.Logger, middleware.Cors, middleware.Preflight)

	wrappedMux := manager.WrapMux(mux)

	server.userHandler.RegisterRoutes(mux, manager)
	server.productHandler.RegisterRoutes(mux, manager)

	port := ":" + fmt.Sprint(server.cnf.HTTPPort)
	fmt.Println("Server started at localhost" + port)
	err := http.ListenAndServe(port, wrappedMux)
	if err != nil {
		fmt.Println("Failed to start server: ", err)
	}
}