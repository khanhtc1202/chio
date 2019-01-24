package src_test

import (
	"testing"

	"github.com/khanhtc1202/chio/src"
	"github.com/stretchr/testify/assert"
)

func TestSourceFile_GetDirPath(t *testing.T) {
	file := src.NewSourceFile("/home/user/go/src/sample/main.go")
	expected := "/home/user/go/src/sample/"

	dir := file.DirPath()
	assert.Equal(t, expected, dir)
}

func TestSourceFile_Name(t *testing.T) {
	file := src.NewSourceFile("/home/user/go/src/sample/main.go")
	expected := "main.go"

	fileName := file.Name()
	assert.Equal(t, expected, fileName)
}
