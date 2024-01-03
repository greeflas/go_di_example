package handler

import (
	"io"
	"log"
	"net/http"
)

type EchoHandler struct {
	logger *log.Logger
}

func NewEchoHandler(logger *log.Logger) *EchoHandler {
	return &EchoHandler{
		logger: logger,
	}
}

func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (h *EchoHandler) Pattern() string {
	return "/echo"
}
