package cmd

import (
	"ecommerce/config"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
)

func Serve() {
	cnf := config.GetConfig()

	userRepo := repo.NewUserRepo()
	productRepo := repo.NewProductRepo()
	
	userHandler := user.NewHandler(userRepo)
	productHandler := product.NewHandler(productRepo)

	server := rest.NewServer(cnf, userHandler, productHandler)
	server.Start()
}