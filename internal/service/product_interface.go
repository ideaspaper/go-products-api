package service

import (
	"productsapi/internal/dto/request"
	"productsapi/internal/dto/response"
)

type IProduct interface {
	AddProduct(productDTO *request.ProductInsert) (*response.Product, error)
	GetProducts(page, size int, search string, orderDir bool) ([]*response.Product, error)
	GetProductById(id int) (*response.Product, error)
	UpdateProduct(id int, productDTO *request.ProductUpdate) (*response.Product, error)
	UpdateProductQuantity(id int, productDTO *request.ProductUpdateQuantity) (*response.Product, error)
	DeleteProduct(id int) (*response.Product, error)
}
