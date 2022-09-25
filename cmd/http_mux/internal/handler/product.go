package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"productsapi/internal/dto/request"
	"productsapi/internal/repository"
	"strconv"

	"github.com/gorilla/mux"
)

func (h Handler) AddProduct(w http.ResponseWriter, r *http.Request) {
	var productDTO *request.ProductInsert
	if err := json.NewDecoder(r.Body).Decode(&productDTO); err != nil {
		log.Printf("%s: %v\n", CtxAddProduct, err)
		SendResponse(
			w,
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			nil,
		)
		return
	}
	if err := h.validator.Struct(productDTO); err != nil {
		log.Printf("%s: %v\n", CtxAddProduct, err)
		SendResponse(
			w,
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			nil,
		)
		return
	}
	product, err := h.productService.AddProduct(productDTO)
	if err != nil {
		log.Printf("%s: %v\n", CtxAddProduct, err)
		SendResponse(
			w,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			nil,
		)
		return
	}
	SendResponse(
		w,
		http.StatusCreated,
		http.StatusText(http.StatusCreated),
		product,
	)
}

func (h Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, _, err := h.productService.GetProducts(0, 0, "", false)
	if err != nil {
		log.Printf("%s: %v\n", CtxGetProducts, err)
		SendResponse(
			w,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			nil,
		)
		return
	}
	SendResponse(
		w,
		http.StatusOK,
		http.StatusText(http.StatusOK),
		products,
	)
}

func (h Handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("%s: %v\n", CtxGetProductById, err)
		SendResponse(
			w,
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			nil,
		)
		return
	}
	product, err := h.productService.GetProductById(id)
	if errors.Is(err, &repository.ErrDataNotFound) {
		log.Printf("%s: %v\n", CtxGetProductById, err)
		SendResponse(
			w,
			http.StatusNotFound,
			http.StatusText(http.StatusNotFound),
			nil,
		)
		return
	}
	if err != nil {
		log.Printf("%s: %v\n", CtxGetProductById, err)
		SendResponse(
			w,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			nil,
		)
		return
	}
	SendResponse(
		w,
		http.StatusOK,
		http.StatusText(http.StatusOK),
		product,
	)
}

func (h Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("%s: %v\n", CtxUpdateProduct, err)
		SendResponse(
			w,
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			nil,
		)
		return
	}
	var productDTO *request.ProductUpdate
	if err := json.NewDecoder(r.Body).Decode(&productDTO); err != nil {
		log.Printf("%s: %v\n", CtxUpdateProduct, err)
		SendResponse(
			w,
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			nil,
		)
		return
	}
	if err := h.validator.Struct(productDTO); err != nil {
		log.Printf("%s: %v\n", CtxUpdateProduct, err)
		SendResponse(
			w,
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			nil,
		)
		return
	}
	product, err := h.productService.UpdateProduct(id, productDTO)
	if errors.Is(err, &repository.ErrDataNotFound) {
		log.Printf("%s: %v\n", CtxUpdateProduct, err)
		SendResponse(
			w,
			http.StatusNotFound,
			http.StatusText(http.StatusNotFound),
			nil,
		)
		return
	}
	if err != nil {
		log.Printf("%s: %v\n", CtxUpdateProduct, err)
		SendResponse(
			w,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			nil,
		)
		return
	}
	SendResponse(
		w,
		http.StatusOK,
		http.StatusText(http.StatusOK),
		product,
	)
}

func (h Handler) UpdateProductQuantity(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("%s: %v\n", CtxUpdateProductQuantity, err)
		SendResponse(
			w,
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			nil,
		)
		return
	}
	var productDTO *request.ProductUpdateQuantity
	if err := json.NewDecoder(r.Body).Decode(&productDTO); err != nil {
		log.Printf("%s: %v\n", CtxUpdateProductQuantity, err)
		SendResponse(
			w,
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			nil,
		)
		return
	}
	if err := h.validator.Struct(productDTO); err != nil {
		log.Printf("%s: %v\n", CtxUpdateProductQuantity, err)
		SendResponse(
			w,
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			nil,
		)
		return
	}
	product, err := h.productService.UpdateProductQuantity(id, productDTO)
	if errors.Is(err, &repository.ErrDataNotFound) {
		log.Printf("%s: %v\n", CtxUpdateProductQuantity, err)
		SendResponse(
			w,
			http.StatusNotFound,
			http.StatusText(http.StatusNotFound),
			nil,
		)
		return
	}
	if err != nil {
		log.Printf("%s: %v\n", CtxUpdateProductQuantity, err)
		SendResponse(
			w,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			nil,
		)
		return
	}
	SendResponse(
		w,
		http.StatusOK,
		http.StatusText(http.StatusOK),
		product,
	)
}

func (h Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("%s: %v\n", CtxDeleteProduct, err)
		SendResponse(
			w,
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			nil,
		)
		return
	}
	product, err := h.productService.DeleteProduct(id)
	if errors.Is(err, &repository.ErrDataNotFound) {
		log.Printf("%s: %v\n", CtxDeleteProduct, err)
		SendResponse(
			w,
			http.StatusNotFound,
			http.StatusText(http.StatusNotFound),
			nil,
		)
		return
	}
	if err != nil {
		log.Printf("%s: %v\n", CtxDeleteProduct, err)
		SendResponse(
			w,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			nil,
		)
		return
	}
	SendResponse(
		w,
		http.StatusOK,
		http.StatusText(http.StatusOK),
		product,
	)
}
