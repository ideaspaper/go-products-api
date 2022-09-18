package pg

import (
	"context"
	"errors"
	"fmt"
	"productsapi/internal/dto/request"
	"productsapi/internal/model"
	"productsapi/internal/repository"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductConfig struct {
	Db *pgxpool.Pool
}

type product struct {
	db *pgxpool.Pool
}

func NewProduct(config *ProductConfig) repository.IProduct {
	return &product{
		db: config.Db,
	}
}

func (p *product) AddProduct(productDTO *request.ProductInsert) (*model.Product, error) {
	product := &model.Product{}
	err := p.db.QueryRow(
		context.Background(),
		`
			INSERT INTO products_tab (name, description, quantity)
			VALUES ($1, $2, $3)
			RETURNING "ID", name, description, quantity;
		`,
		productDTO.Name,
		productDTO.Description,
		productDTO.Quantity,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Quantity,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"%s: %w",
			repository.CtxAddProduct,
			repository.ErrUnknown.SetError(err),
		)
	}
	return product, nil
}

func (p *product) GetProducts(page, size int, search string, orderDir bool) ([]*model.Product, error) {
	products := []*model.Product{}
	offset := (page - 1) * size
	limit := size
	rows, err := p.db.Query(
		context.Background(),
		`
			SELECT "ID", name, description, quantity
			FROM products_tab
			WHERE name ILIKE $1
			ORDER BY
				CASE WHEN $2 = true then "ID" END DESC,
				CASE WHEN $2 = false then "ID" END ASC
			OFFSET $3
			LIMIT $4
		`,
		"%"+search+"%",
		orderDir,
		offset,
		limit,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"%s: %w",
			repository.CtxGetProducts,
			repository.ErrUnknown.SetError(err),
		)
	}
	for rows.Next() {
		product := &model.Product{}
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Quantity,
		); err != nil {
			return nil, fmt.Errorf(
				"%s: %w",
				repository.CtxGetProducts,
				repository.ErrUnknown.SetError(err),
			)
		}
		products = append(products, product)
	}
	return products, nil
}

func (p *product) GetProductById(id int) (*model.Product, error) {
	product := &model.Product{}
	err := p.db.QueryRow(
		context.Background(),
		`
			SELECT "ID", name, description, quantity
			FROM products_tab
			WHERE "ID" = $1;
		`,
		id,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Quantity,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf(
			"%s: %w",
			repository.CtxGetProductById,
			repository.ErrDataNotFound.SetError(err),
		)
	}
	if err != nil {
		return nil, fmt.Errorf(
			"%s: %w",
			repository.CtxGetProductById,
			repository.ErrUnknown.SetError(err),
		)
	}
	return product, nil
}

func (p *product) UpdateProduct(id int, productDTO *request.ProductUpdate) (*model.Product, error) {
	product := &model.Product{}
	err := p.db.QueryRow(
		context.Background(),
		`
			UPDATE products_tab
			SET name = $1, description = $2
			WHERE "ID" = $3
			RETURNING "ID", name, description, quantity;
		`,
		productDTO.Name,
		productDTO.Description,
		id,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Quantity,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf(
			"%s: %w",
			repository.CtxUpdateProduct,
			repository.ErrDataNotFound.SetError(err),
		)
	}
	if err != nil {
		return nil, fmt.Errorf(
			"%s: %w",
			repository.CtxUpdateProduct,
			repository.ErrUnknown.SetError(err),
		)
	}
	return product, nil
}

func (p *product) UpdateProductQuantity(id int, productDTO *request.ProductUpdateQuantity) (*model.Product, error) {
	product := &model.Product{}
	err := p.db.QueryRow(
		context.Background(),
		`
			UPDATE products_tab
			SET quantity = $1
			WHERE "ID" = $2
			RETURNING "ID", name, description, quantity;
		`,
		productDTO.Quantity,
		id,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Quantity,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf(
			"%s: %w",
			repository.CtxUpdateProductQuantity,
			repository.ErrDataNotFound.SetError(err),
		)
	}
	if err != nil {
		return nil, fmt.Errorf(
			"%s: %w",
			repository.CtxUpdateProductQuantity,
			repository.ErrUnknown.SetError(err),
		)
	}
	return product, nil
}

func (p *product) DeleteProduct(id int) (*model.Product, error) {
	product := &model.Product{}
	err := p.db.QueryRow(
		context.Background(),
		`
			DELETE FROM products_tab
			WHERE "ID" = $1
			RETURNING "ID", name, description, quantity;
		`,
		id,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Quantity,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf(
			"%s: %w",
			repository.CtxDeleteProduct,
			repository.ErrDataNotFound.SetError(err),
		)
	}
	if err != nil {
		return nil, fmt.Errorf(
			"%s: %w",
			repository.CtxDeleteProduct,
			repository.ErrUnknown.SetError(err),
		)
	}
	return nil, nil
}
