package main

import (
	"github.com/greeflas/go_di_example/pkg/di"
	"testing"
)

func TestRunCommand(t *testing.T) {
	if err := runCommand(di.EnvTest); err != nil {
		t.Fatal(err)
	}
}
