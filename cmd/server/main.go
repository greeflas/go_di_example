package main

import (
	"log"

	"github.com/greeflas/go_di_example/internal/handler"
	"github.com/greeflas/go_di_example/pkg/di"
	"github.com/greeflas/go_di_example/pkg/server"
	"go.uber.org/dig"
)

func main() {
	if err := runApp(di.EnvProd); err != nil {
		log.Fatal(err)
	}
}

func runApp(env di.Environment) error {
	c := di.BuildContainer(env)
	c = addAppSpecificDependencies(c)

	return c.Invoke(func(apiServer *server.APIServer) error {
		return apiServer.Start()
	})
}

func addAppSpecificDependencies(container *dig.Container) *dig.Container {
	container.Provide(handler.NewHelloHandler, dig.As(new(server.Handler)), dig.Group("handlers"))
	container.Provide(handler.NewEchoHandler, dig.As(new(server.Handler)), dig.Group("handlers"))
	container.Provide(server.NewAPIServer)

	container.Decorate(func(logger *log.Logger) *log.Logger {
		logger.SetPrefix("[server] ")

		return logger
	})

	return container
}
