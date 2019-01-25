package pkg_test

import (
	"fmt"
	"testing"

	"github.com/khanhtc1202/chio/pkg"
	"github.com/stretchr/testify/assert"
)

func TestModuleFactory_DirectoryAsModuleFileLevel(t *testing.T) {
	rootPath := "../sample"

	moduleFac := pkg.NewModuleFactory()
	modules, err := moduleFac.DirectoryAsModule(rootPath, pkg.GO)

	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(modules))

	for _, module := range modules {
		files := module.GetSourceFilesPath()
		fmt.Println(files)
	}
}

func TestModuleFactory_DirectoryAsModuleNoSubDir(t *testing.T) {
	rootPath := "../sample/sub"

	moduleFac := pkg.NewModuleFactory()
	modules, err := moduleFac.DirectoryAsModule(rootPath, pkg.GO)

	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(modules))

	for _, module := range modules {
		files := module.GetSourceFilesPath()
		fmt.Println(files)
	}
}

func TestModuleFactory_DirectoryAsModule_NotExistedDir(t *testing.T) {
	rootPath := "./fff"

	moduleFac := pkg.NewModuleFactory()
	_, err := moduleFac.DirectoryAsModule(rootPath, pkg.GO)

	assert.NotNil(t, err)
}
