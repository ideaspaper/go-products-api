package handler

import (
	"encoding/json"
	"net/http"
	"productsapi/internal/dto"
	"productsapi/internal/service"

	"github.com/go-playground/validator/v10"
)

type HandlerConfig struct {
	ProductService service.IProduct
	Validator      *validator.Validate
}

type Handler struct {
	productService service.IProduct
	validator      *validator.Validate
}

func New(config *HandlerConfig) *Handler {
	return &Handler{
		productService: config.ProductService,
		validator:      config.Validator,
	}
}

func SendResponse(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(
		&dto.Standard{
			Code:    code,
			Message: message,
			Data:    data,
		},
	)
}
