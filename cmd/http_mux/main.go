package main

import (
	"log"
	"net/http"
	"productsapi/cmd/http_mux/internal/handler"
	"productsapi/cmd/http_mux/internal/router"
	"productsapi/internal/model"
	"productsapi/internal/repository/memory"
	"productsapi/internal/service"

	"github.com/go-playground/validator/v10"
)

var products = []*model.Product{
	{
		ID:          1,
		Name:        "Latitude 5420",
		Description: "A generously equipped allround business laptop that leaves nothing to be desired in terms of connectivity and security.",
		Quantity:    10,
	},
	{
		ID:          2,
		Name:        "ULTRABOOST 5 DNA RUNNING SPORTSWEAR LIFESTYLE SHOES",
		Description: "EVERYDAY RUNNING SHOES MADE IN PART WITH PARLEY OCEAN PLASTIC.",
		Quantity:    23,
	},
	{
		ID:          3,
		Name:        "D.O.N. ISSUE #3 SHOES",
		Description: "SIGNATURE SHOES FROM ADIDAS BASKETBALL AND DONOVAN MITCHELL.",
		Quantity:    10,
	},
}

func initHandlers() *handler.Handler {
	productRepositoryConfig := &memory.ProductConfig{Db: products}
	productRepository := memory.NewProduct(productRepositoryConfig)
	productServiceConfig := &service.ProductConfig{ProductRepository: productRepository}
	productService := service.NewProduct(productServiceConfig)
	handlerConfig := &handler.HandlerConfig{
		ProductService: productService,
		Validator:      validator.New(),
	}
	return handler.New(handlerConfig)
}

func main() {
	h := initHandlers()
	r := router.New(h)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalln(err)
	}
}
