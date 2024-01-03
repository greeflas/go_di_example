package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/greeflas/go_di_example/internal/service"
)

type HelloHandler struct {
	logger          *log.Logger
	greetingService *service.GreetingService
}

func NewHelloHandler(
	logger *log.Logger,
	greetingService *service.GreetingService,
) *HelloHandler {
	return &HelloHandler{
		logger:          logger,
		greetingService: greetingService,
	}
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.URL.Query().Get("userId"))
	if err != nil {
		h.logger.Println(err)
	}

	message, err := h.greetingService.GetGreetingMessage(userId)
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}

	if _, err := fmt.Fprint(w, message); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (h *HelloHandler) Pattern() string {
	return "/hello"
}
