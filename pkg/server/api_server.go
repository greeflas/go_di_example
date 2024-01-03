package server

import (
	"log"
	"net/http"

	"github.com/greeflas/go_di_example/internal/handler"
)

type APIServer struct {
	logger     *log.Logger
	httpServer *http.Server
}

func NewAPIServer(
	logger *log.Logger,
	helloHandler *handler.HelloHandler,
	echoHandler *handler.EchoHandler,
) *APIServer {
	mux := http.NewServeMux()
	mux.Handle(helloHandler.Pattern(), helloHandler)
	mux.Handle(echoHandler.Pattern(), echoHandler)

	return &APIServer{
		logger: logger,
		httpServer: &http.Server{
			Addr:    ":8080",
			Handler: mux,
		},
	}
}

func (s *APIServer) Start() error {
	s.logger.Println("Starting server...")

	return s.httpServer.ListenAndServe()
}
