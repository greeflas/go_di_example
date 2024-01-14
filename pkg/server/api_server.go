package server

import (
	"log"
	"net/http"

	"go.uber.org/dig"
)

type Handler interface {
	http.Handler
	Pattern() string
}

type APIServerParams struct {
	dig.In

	Handlers []Handler `group:"handlers"`
}

type APIServer struct {
	logger     *log.Logger
	httpServer *http.Server
}

func NewAPIServer(
	logger *log.Logger,
	params APIServerParams,
) *APIServer {
	mux := http.NewServeMux()
	for _, handler := range params.Handlers {
		mux.Handle(handler.Pattern(), handler)
	}

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
