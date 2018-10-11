package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModule_AddSourceFile_FirstTimeAddFileToEmptyModule(t *testing.T) {
	emptyModule := NewModule(GO)
	file := NewSourceFile("/home/user/go/src/sample/main.go")

	err := emptyModule.AddSourceFile(file)
	assert.Nil(t, err)
	assert.Equal(t, "/home/user/go/src/sample/", emptyModule.RootPath)
}

func TestModule_AddSourceFile_AddFileToExistedModule(t *testing.T) {
	emptyModule := NewModule(GO)
	file1 := NewSourceFile("/home/user/go/src/sample/main.go")
	file2 := NewSourceFile("/home/user/go/src/sample/submodule/sample.go")

	err := emptyModule.AddSourceFile(file1)
	err = emptyModule.AddSourceFile(file2)
	assert.Nil(t, err)
	assert.Equal(t, "/home/user/go/src/sample/", emptyModule.RootPath)
}
