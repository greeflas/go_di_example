package main

import (
	"log"
	"os"

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
	serverScope := addAppSpecificDependencies(c)

	if err := dig.Visualize(c, os.Stdout); err != nil {
		return err
	}

	err := serverScope.Invoke(func(apiServer *server.APIServer) error {
		return apiServer.Start()
	})

	return err
}

func addAppSpecificDependencies(container *dig.Container) *dig.Scope {
	container.Provide(handler.NewHelloHandler, dig.As(new(server.Handler)), dig.Group("handlers"))
	container.Provide(handler.NewEchoHandler, dig.As(new(server.Handler)), dig.Group("handlers"))

	container.Decorate(func(logger *log.Logger) *log.Logger {
		logger.SetPrefix("[server] ")

		return logger
	})

	serverScope := container.Scope("api_server")
	serverScope.Provide(server.NewAPIServer)

	return serverScope
}
