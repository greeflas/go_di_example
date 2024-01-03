package main

import (
	"testing"

	"github.com/greeflas/go_di_example/pkg/di"
)

func TestRunApp(t *testing.T) {
	if err := runApp(di.EnvTest); err != nil {
		t.Fatal(err)
	}
}
