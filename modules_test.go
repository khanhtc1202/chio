package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModules_Add(t *testing.T) {
	modules := NewModules()
	module := NewModule(GO)
	module.RootPath = "/"

	err := modules.Add(module)
	assert.Nil(t, err)
}
