package src_test

import (
	"testing"

	"github.com/khanhtc1202/chio/src"
	"github.com/stretchr/testify/assert"
)

func TestModule_AddSourceFile_FirstTimeAddFileToEmptyModule(t *testing.T) {
	emptyModule := src.NewModule("/home/user/go/src/sample/", nil)
	file := src.NewSourceFile("/home/user/go/src/sample/main.go")

	err := emptyModule.AddSourceFile(file)
	assert.Nil(t, err)
	assert.Equal(t, "/home/user/go/src/sample/", emptyModule.RootPath)
}

func TestModule_AddSourceFile_AddFileToExistedModule(t *testing.T) {
	emptyModule := src.NewModule("/home/user/go/src/sample/", nil)
	file1 := src.NewSourceFile("/home/user/go/src/sample/main.go")
	file2 := src.NewSourceFile("/home/user/go/src/sample/submodule/sample.go")

	err := emptyModule.AddSourceFile(file1)
	err = emptyModule.AddSourceFile(file2)
	assert.Nil(t, err)
	assert.Equal(t, "/home/user/go/src/sample/", emptyModule.RootPath)
}

func TestModule_GetSourceFilesPath(t *testing.T) {
	emptyModule := src.NewModule("/home/user/go/src/sample/", nil)
	file1 := src.NewSourceFile("/home/user/go/src/sample/main.go")
	_ = emptyModule.AddSourceFile(file1)

	filesPath := emptyModule.GetSourceFilesPath()
	assert.Equal(t, "/home/user/go/src/sample/main.go", filesPath[0])
}
