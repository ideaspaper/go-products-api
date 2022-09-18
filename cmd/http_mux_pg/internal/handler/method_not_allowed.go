package handler

import (
	"net/http"
)

func (h Handler) MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	SendResponse(
		w,
		http.StatusMethodNotAllowed,
		http.StatusText(http.StatusMethodNotAllowed),
		nil,
	)
}
