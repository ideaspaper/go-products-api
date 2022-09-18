package service

import (
	"fmt"
	"productsapi/internal/dto/request"
	"productsapi/internal/dto/response"
	"productsapi/internal/repository"
)

type ProductConfig struct {
	ProductRepository repository.IProduct
}

type product struct {
	productRepository repository.IProduct
}

func NewProduct(config *ProductConfig) IProduct {
	return product{
		productRepository: config.ProductRepository,
	}
}

func (p product) AddProduct(productDTO *request.ProductInsert) (*response.Product, error) {
	product, err := p.productRepository.AddProduct(productDTO)
	if err != nil {
		return nil, fmt.Errorf(
			"%s: %w",
			CtxAddProduct,
			err,
		)
	}
	return product.ToDTO(), nil
}

func (p product) GetProducts(page, size int, search string, orderDir bool) ([]*response.Product, error) {
	products, err := p.productRepository.GetProducts(page, size, search, orderDir)
	if err != nil {
		return nil, fmt.Errorf(
			"%s: %w",
			CtxGetProducts,
			err,
		)
	}
	result := []*response.Product{}
	for _, product := range products {
		result = append(result, product.ToDTO())
	}
	return result, nil
}

func (p product) GetProductById(id int) (*response.Product, error) {
	product, err := p.productRepository.GetProductById(id)
	if err != nil {
		return nil, fmt.Errorf(
			"%s: %w",
			CtxGetProductById,
			err,
		)
	}
	return product.ToDTO(), nil
}

func (p product) UpdateProduct(id int, productDTO *request.ProductUpdate) (*response.Product, error) {
	product, err := p.productRepository.UpdateProduct(id, productDTO)
	if err != nil {
		return nil, fmt.Errorf(
			"%s: %w",
			CtxUpdateProduct,
			err,
		)
	}
	return product.ToDTO(), nil
}

func (p product) UpdateProductQuantity(id int, productDTO *request.ProductUpdateQuantity) (*response.Product, error) {
	product, err := p.productRepository.UpdateProductQuantity(id, productDTO)
	if err != nil {
		return nil, fmt.Errorf(
			"%s: %w",
			CtxUpdateProductQuantity,
			err,
		)
	}
	return product.ToDTO(), nil
}

func (p product) DeleteProduct(id int) (*response.Product, error) {
	product, err := p.productRepository.DeleteProduct(id)
	if err != nil {
		return nil, fmt.Errorf(
			"%s: %w",
			CtxDeleteProduct,
			err,
		)
	}
	return product.ToDTO(), nil
}
