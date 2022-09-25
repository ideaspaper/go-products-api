package memory

import (
	"fmt"
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

func (p *product) GetProducts(page, size int, search string, orderDir bool) ([]*model.Product, int, error) {
	return p.db, 0, nil
}

func (p *product) GetProductById(id int) (*model.Product, error) {
	for _, v := range p.db {
		if v.ID == id {
			return v, nil
		}
	}
	return nil, fmt.Errorf(
		"%s: %w",
		repository.CtxGetProductById,
		&repository.ErrDataNotFound,
	)
}

func (p *product) UpdateProduct(id int, productDTO *request.ProductUpdate) (*model.Product, error) {
	var pr *model.Product
	for _, v := range p.db {
		if v.ID == id {
			pr = v
		}
	}
	if pr == nil {
		return nil, fmt.Errorf(
			"%s: %w",
			repository.CtxUpdateProduct,
			&repository.ErrDataNotFound,
		)
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
		return nil, fmt.Errorf(
			"%s: %w",
			repository.CtxUpdateProductQuantity,
			&repository.ErrDataNotFound,
		)
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
		return nil, fmt.Errorf(
			"%s: %w",
			repository.CtxDeleteProduct,
			&repository.ErrDataNotFound,
		)
	}
	return pr, nil
}
