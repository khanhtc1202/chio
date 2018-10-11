package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModuleFactory_DirectoryAsModuleFileLevel(t *testing.T) {
	rootPath := "/Users/khanh.tran/workspace/go/src/github.com/khanhtc1202/chio/sample"

	moduleFac := NewModuleFactory()
	modules, err := moduleFac.DirectoryAsModule(rootPath, GO)

	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(modules))

	for _, module := range modules {
		files := module.GetSourceFilesPath()
		fmt.Println(files)
	}
}
