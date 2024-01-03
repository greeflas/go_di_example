package di

import (
	"log"

	"github.com/greeflas/go_di_example/internal/repository"
	"github.com/greeflas/go_di_example/internal/service"
	"github.com/greeflas/go_di_example/pkg/db"
	"go.uber.org/dig"
)

type Environment int

const (
	EnvProd Environment = iota
	EnvTest
)

func BuildContainer(env Environment) *dig.Container {
	dryRun := false
	if env == EnvTest {
		dryRun = true
	}

	container := dig.New(dig.DryRun(dryRun))

	container.Provide(func() *log.Logger {
		return log.Default()
	})
	container.Provide(db.NewConnection)
	container.Provide(repository.NewUserRepository)
	container.Provide(repository.NewMessageRepository)
	container.Provide(service.NewGreetingService)

	return container
}
