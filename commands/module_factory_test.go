package commands

import (
	"fmt"
	"testing"

	"github.com/khanhtc1202/chio/entity"
)

var rootPath = "/Users/khanh.tran/workspace/go/src/github.com/khanhtc1202/chio"

func TestModuleFactory_DirectoryAsModuleFileLevel(t *testing.T) {
	moduleFac := NewModuleFactory()
	modules, err := moduleFac.DirectoryAsModule(rootPath, entity.GO)
	if err != nil {
		t.Fatal("Error on load file level test:", err)
	}
	if len(*modules) == 0 {
		t.Fatal("Count false in load module file level")
	}
	for _, module := range *modules {
		files := module.GetSourceFilesPath()
		fmt.Println(files)
	}
}
