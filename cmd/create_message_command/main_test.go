package main

import (
	"testing"

	"github.com/greeflas/go_di_example/pkg/di"
)

func TestRunCommand(t *testing.T) {
	if err := runCommand(di.EnvTest); err != nil {
		t.Fatal(err)
	}
}
