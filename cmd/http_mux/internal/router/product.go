package router

import (
	"productsapi/cmd/http_mux/internal/handler"

	"github.com/gorilla/mux"
)

func InitProduct(r *mux.Router, h *handler.Handler) {
	s := r.PathPrefix("/products").Subrouter()
	s.HandleFunc("", h.AddProduct).Methods("POST")
	s.HandleFunc("", h.GetProducts).Methods("GET")
	s.HandleFunc("/{id}", h.GetProductById).Methods("GET")
	s.HandleFunc("/{id}", h.UpdateProduct).Methods("PUT")
	s.HandleFunc("/{id}", h.UpdateProductQuantity).Methods("PATCH")
	s.HandleFunc("/{id}", h.DeleteProduct).Methods("DELETE")
}
