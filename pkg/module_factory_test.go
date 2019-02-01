package pkg_test

import (
	"fmt"
	"testing"

	"github.com/khanhtc1202/chio/pkg"
	"github.com/stretchr/testify/assert"
)

func TestModuleFactory_NDepthDirectoryAsModule_ComplexModule(t *testing.T) {
	rootPath := "../sample"

	moduleFac := pkg.NewNDepthDirectoryAsModule(rootPath)
	modules, err := moduleFac.Load(pkg.GO)

	assert.Nil(t, err)
	assert.Equal(t, 4, len(modules))

	for _, module := range modules {
		files := module.GetSourceFilesPath()
		fmt.Println(files)
	}
}

func TestModuleFactory_NDepthDirectoryAsModule_NoSubDir(t *testing.T) {
	rootPath := "../sample/sub"

	moduleFac := pkg.NewNDepthDirectoryAsModule(rootPath)
	modules, err := moduleFac.Load(pkg.GO)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(modules))

	for _, module := range modules {
		files := module.GetSourceFilesPath()
		fmt.Println(files)
	}
}

func TestModuleFactory_NDepthDirectoryAsModule_NotExistedDir(t *testing.T) {
	rootPath := "./fff"

	moduleFac := pkg.NewNDepthDirectoryAsModule(rootPath)
	_, err := moduleFac.Load(pkg.GO)

	assert.NotNil(t, err)
}

func TestModuleFactory_OneDepthDirectoryAsModule_ComplexModule(t *testing.T) {
	rootPath := "../sample"

	moduleFac := pkg.NewOneDepthDirectoryAsModule(rootPath)
	modules, err := moduleFac.Load(pkg.GO)

	assert.Nil(t, err)
	assert.Equal(t, 2, len(modules))

	for _, module := range modules {
		files := module.GetSourceFilesPath()
		fmt.Println(files)
	}
}
