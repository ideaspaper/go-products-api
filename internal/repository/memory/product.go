package memory

import (
	"log"
	"productsapi/internal/dto/request"
	"productsapi/internal/model"
	"productsapi/internal/repository"
)

type ProductConfig struct {
	Db []*model.Product
}

type product struct {
	db []*model.Product
}

func NewProduct(config *ProductConfig) repository.IProduct {
	return &product{
		db: config.Db,
	}
}

func (p *product) AddProduct(productDTO *request.ProductInsert) (*model.Product, error) {
	id := 1
	if len(p.db) != 0 {
		id = p.db[len(p.db)-1].ID + 1
	}
	pr := &model.Product{
		ID:          id,
		Name:        productDTO.Name,
		Description: productDTO.Description,
		Quantity:    *productDTO.Quantity,
	}
	p.db = append(p.db, pr)
	return pr, nil
}

func (p *product) GetProducts(page, size int, search, sortBy, sortDir string) ([]*model.Product, error) {
	return p.db, nil
}

func (p *product) GetProductById(id int) (*model.Product, error) {
	for _, v := range p.db {
		if v.ID == id {
			return v, nil
		}
	}
	return nil, repository.ErrRepositoryDataNotFound
}

func (p *product) UpdateProduct(id int, productDTO *request.ProductUpdate) (*model.Product, error) {
	var pr *model.Product
	for _, v := range p.db {
		if v.ID == id {
			pr = v
		}
	}
	if pr == nil {
		log.Println(repository.LogUpdateProduct, repository.ErrRepositoryDataNotFound)
		return nil, repository.ErrRepositoryDataNotFound
	}
	pr.Name = productDTO.Name
	pr.Description = productDTO.Description
	return pr, nil
}

func (p *product) UpdateProductQuantity(id int, productDTO *request.ProductUpdateQuantity) (*model.Product, error) {
	var pr *model.Product
	for _, v := range p.db {
		if v.ID == id {
			pr = v
		}
	}
	if pr == nil {
		log.Println(repository.LogUpdateProductQuantity, repository.ErrRepositoryDataNotFound)
		return nil, repository.ErrRepositoryDataNotFound
	}
	pr.Quantity = *productDTO.Quantity
	return pr, nil
}

func (p *product) DeleteProduct(id int) (*model.Product, error) {
	var pr *model.Product
	for i, v := range p.db {
		if v.ID == id {
			pr = v
			p.db = append(p.db[:i], p.db[i+1:]...)
		}
	}
	if pr == nil {
		log.Println(repository.LogDeleteProduct, repository.ErrRepositoryDataNotFound)
		return nil, repository.ErrRepositoryDataNotFound
	}
	return pr, nil
}
