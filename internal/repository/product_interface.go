package repository

import (
	"productsapi/internal/dto/request"
	"productsapi/internal/model"
)

type IProduct interface {
	AddProduct(productDTO *request.ProductInsert) (*model.Product, error)
	GetProducts(page, size int, search string, orderDir bool) ([]*model.Product, error)
	GetProductById(id int) (*model.Product, error)
	UpdateProduct(id int, productDTO *request.ProductUpdate) (*model.Product, error)
	UpdateProductQuantity(id int, productDTO *request.ProductUpdateQuantity) (*model.Product, error)
	DeleteProduct(id int) (*model.Product, error)
}
