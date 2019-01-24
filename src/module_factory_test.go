package src_test

import (
	"fmt"
	"testing"

	"github.com/khanhtc1202/chio/src"
	"github.com/stretchr/testify/assert"
)

func TestModuleFactory_DirectoryAsModuleFileLevel(t *testing.T) {
	rootPath := "../sample"

	moduleFac := src.NewModuleFactory()
	modules, err := moduleFac.DirectoryAsModule(rootPath, src.GO)

	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(modules))

	for _, module := range modules {
		files := module.GetSourceFilesPath()
		fmt.Println(files)
	}
}

func TestModuleFactory_DirectoryAsModuleNoSubDir(t *testing.T) {
	rootPath := "../sample/sub"

	moduleFac := src.NewModuleFactory()
	modules, err := moduleFac.DirectoryAsModule(rootPath, src.GO)

	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(modules))

	for _, module := range modules {
		files := module.GetSourceFilesPath()
		fmt.Println(files)
	}
}
