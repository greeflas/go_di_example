package main

import (
	"github.com/greeflas/go_di_example/pkg/di"
	"testing"
)

func TestRunApp(t *testing.T) {
	if err := runApp(di.EnvTest); err != nil {
		t.Fatal(err)
	}
}
