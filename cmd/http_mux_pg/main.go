package main

import (
	"log"
	"net/http"
	"productsapi/cmd/http_mux_pg/internal/db"
	"productsapi/cmd/http_mux_pg/internal/handler"
	"productsapi/cmd/http_mux_pg/internal/router"
	"productsapi/internal/repository/pg"
	"productsapi/internal/service"

	"github.com/go-playground/validator/v10"
)

func initHandlers() *handler.Handler {
	productRepositoryConfig := &pg.ProductConfig{Db: db.Get()}
	productRepository := pg.NewProduct(productRepositoryConfig)
	productServiceConfig := &service.ProductConfig{ProductRepository: productRepository}
	productService := service.NewProduct(productServiceConfig)
	handlerConfig := &handler.HandlerConfig{
		ProductService: productService,
		Validator:      validator.New(),
	}
	return handler.New(handlerConfig)
}

func main() {
	if err := db.Connect(); err != nil {
		log.Fatalln(err)
	}
	h := initHandlers()
	r := router.New(h)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalln(err)
	}
}
