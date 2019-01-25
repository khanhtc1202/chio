package pkg_test

import (
	"testing"

	"github.com/khanhtc1202/chio/pkg"
	"github.com/stretchr/testify/assert"
)

func TestSourceFile_GetDirPath(t *testing.T) {
	file := pkg.NewSourceFile("/home/user/go/src/sample/main.go")
	expected := "/home/user/go/src/sample/"

	dir := file.DirPath()
	assert.Equal(t, expected, dir)
}

func TestSourceFile_GetAbsoluteDirPath(t *testing.T) {
	file := pkg.NewSourceFile("../sample/main.go")
	expected := "/home/user/go/src/sample/"

	dir := file.DirPath()
	assert.NotEqual(t, expected, dir)
}

func TestSourceFile_Name(t *testing.T) {
	file := pkg.NewSourceFile("/home/user/go/src/sample/main.go")
	expected := "main.go"

	fileName := file.Name()
	assert.Equal(t, expected, fileName)
}
