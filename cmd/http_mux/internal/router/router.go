package router

import (
	"net/http"
	"productsapi/cmd/http_mux/api"
	"productsapi/cmd/http_mux/internal/handler"
	"productsapi/cmd/http_mux/internal/middleware"

	"github.com/gorilla/mux"
)

func New(h *handler.Handler) *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/docs/").Handler(
		http.StripPrefix(
			"/docs/",
			http.FileServer(http.FS(api.Docs)),
		),
	)
	r.Use(middleware.Logging)
	InitProduct(r, h)
	r.NotFoundHandler = http.HandlerFunc(h.NotFound)
	r.MethodNotAllowedHandler = http.HandlerFunc(h.MethodNotAllowed)
	return r
}
