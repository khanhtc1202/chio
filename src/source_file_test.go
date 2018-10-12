package src_test

import (
	"testing"

	"github.com/khanhtc1202/chio/src"
	"github.com/stretchr/testify/assert"
)

func TestSourceFile_GetDirPath(t *testing.T) {
	file := src.NewSourceFile("/home/user/go/src/sample/main.go")
	expected := "/home/user/go/src/sample/"

	dir := file.GetDirPath()
	assert.Equal(t, expected, dir)
}
