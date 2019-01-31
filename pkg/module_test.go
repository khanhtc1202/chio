package pkg_test

import (
	"testing"

	"github.com/khanhtc1202/chio/pkg"
	"github.com/stretchr/testify/assert"
)

func TestModule_AddSourceFile_FirstTimeAddFileToEmptyModule(t *testing.T) {
	emptyModule := pkg.NewModule("/home/user/go/src/sample/")
	file := pkg.NewSourceFile("/home/user/go/src/sample/main.go")

	err := emptyModule.AddSourceFile(file)
	assert.Nil(t, err)
	assert.Equal(t, "/home/user/go/src/sample", emptyModule.RootPath)
}

func TestModule_AddSourceFile_AddFileToExistedModule(t *testing.T) {
	emptyModule := pkg.NewModule("/home/user/go/src/sample/")
	file1 := pkg.NewSourceFile("/home/user/go/src/sample/main.go")
	file2 := pkg.NewSourceFile("/home/user/go/src/sample/submodule/sample.go")

	err := emptyModule.AddSourceFile(file1)
	err = emptyModule.AddSourceFile(file2)
	assert.Nil(t, err)
	assert.Equal(t, "/home/user/go/src/sample", emptyModule.RootPath)
}

func TestModule_GetSourceFilesPath(t *testing.T) {
	emptyModule := pkg.NewModule("/home/user/go/src/sample/")
	file1 := pkg.NewSourceFile("/home/user/go/src/sample/main.go")
	_ = emptyModule.AddSourceFile(file1)

	filesPath := emptyModule.GetSourceFilesPath()
	assert.Equal(t, "/home/user/go/src/sample/main.go", filesPath[0])
}

func TestModule_AddSourceFiles(t *testing.T) {
	emptyModule := pkg.NewModule("/home/user/go/src/sample/")
	files := pkg.SourceFiles{
		pkg.NewSourceFile("/home/user/go/src/sample/main.go"),
		pkg.NewSourceFile("/home/user/go/src/sample/sample.go"),
		pkg.NewSourceFile("/home/user/go/src/sample/sub_dir/main.go"),
	}

	err := emptyModule.AddSourceFiles(files)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(emptyModule.SourceFiles))
}
