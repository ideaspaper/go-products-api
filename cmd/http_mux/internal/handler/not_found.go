package handler

import (
	"net/http"
)

func (h Handler) NotFound(w http.ResponseWriter, r *http.Request) {
	SendResponse(
		w,
		http.StatusNotFound,
		http.StatusText(http.StatusNotFound),
		nil,
	)
}
