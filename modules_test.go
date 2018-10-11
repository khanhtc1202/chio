package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModules_Add_AddNotEmptyModule(t *testing.T) {
	modules := NewModules()
	module := NewModule(nil)
	module.RootPath = "/"

	err := modules.Add(module)
	assert.Nil(t, err)
}

func TestModules_Add_AddEmptyModule(t *testing.T) {
	modules := NewModules()
	module := NewModule(nil)

	err := modules.Add(module)
	assert.NotNil(t, err)
}
