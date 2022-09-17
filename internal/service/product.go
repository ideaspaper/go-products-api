package service

import (
	"log"
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
		log.Println(LogAddProduct, err)
		return nil, err
	}
	return product.ToDTO(), nil
}

func (p product) GetProducts(page, size int, search, orderBy, orderDir string) ([]*response.Product, error) {
	products, err := p.productRepository.GetProducts(page, size, search, orderBy, orderDir)
	if err != nil {
		log.Println(LogGetProducts, err)
		return nil, err
	}
	result := []*response.Product{}
	for _, product := range products {
		result = append(result, product.ToDTO())
	}
	return result, nil
}

func (p product) GetProductById(id int) (*response.Product, error) {
	product, err := p.productRepository.GetProductById(id)
	if err == repository.ErrRepositoryDataNotFound {
		log.Println(LogGetProductById, err)
		return nil, ErrServiceDataNotFound
	}
	if err != nil {
		log.Println(LogGetProductById, err)
		return nil, err
	}
	return product.ToDTO(), nil
}

func (p product) UpdateProduct(id int, productDTO *request.ProductUpdate) (*response.Product, error) {
	product, err := p.productRepository.UpdateProduct(id, productDTO)
	if err == repository.ErrRepositoryDataNotFound {
		log.Println(LogUpdateProduct, err)
		return nil, ErrServiceDataNotFound
	}
	if err != nil {
		log.Println(LogUpdateProduct, err)
		return nil, err
	}
	return product.ToDTO(), nil
}

func (p product) UpdateProductQuantity(id int, productDTO *request.ProductUpdateQuantity) (*response.Product, error) {
	product, err := p.productRepository.UpdateProductQuantity(id, productDTO)
	if err == repository.ErrRepositoryDataNotFound {
		log.Println(LogUpdateProductQuantity, err)
		return nil, ErrServiceDataNotFound
	}
	if err != nil {
		log.Println(LogUpdateProductQuantity, err)
		return nil, err
	}
	return product.ToDTO(), nil
}

func (p product) DeleteProduct(id int) (*response.Product, error) {
	product, err := p.productRepository.DeleteProduct(id)
	if err == repository.ErrRepositoryDataNotFound {
		log.Println(LogDeleteProduct, err)
		return nil, ErrServiceDataNotFound
	}
	if err != nil {
		log.Println(LogDeleteProduct, err)
		return nil, err
	}
	return product.ToDTO(), nil
}
