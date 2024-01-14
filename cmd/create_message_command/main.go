package main

import (
	"errors"
	"log"
	"os"

	"go.uber.org/dig"

	"github.com/greeflas/go_di_example/internal/model"
	"github.com/greeflas/go_di_example/internal/repository"
	"github.com/greeflas/go_di_example/pkg/di"
)

func main() {
	if err := runCommand(di.EnvProd); err != nil {
		log.Fatal(err)
	}
}

func runCommand(env di.Environment) error {
	c := di.BuildContainer(env)

	c.Decorate(func(logger *log.Logger) *log.Logger {
		logger.SetPrefix("[server] ")

		return logger
	})

	if err := dig.Visualize(c, os.Stdout); err != nil {
		return err
	}

	return c.Invoke(func(messageRepository *repository.MessageRepository) error {
		if len(os.Args) < 2 {
			return errors.New("pattern arg is required")
		}

		pattern := os.Args[1]
		message := model.NewMessage(pattern)

		return messageRepository.Create(message)
	})
}
