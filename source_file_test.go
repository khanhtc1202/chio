package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSourceFile_GetDirPath(t *testing.T) {
	file := NewSourceFile("/home/user/go/src/sample/main.go")
	expected := "/home/user/go/src/sample/"

	dir := file.GetDirPath()
	assert.Equal(t, expected, dir)
}
