package commands_test

import (
	"testing"

	"fmt"

	"github.com/khanhtc1202/chio/commands"
	"github.com/khanhtc1202/chio/domain"
)

var rootPath = "/Users/khanh.tran/workspace/go/src/github.com/khanhtc1202/chio"

func TestModuleFactory_LoadFileLevel(t *testing.T) {
	moduleFac := commands.NewModuleFactory()
	modules, err := moduleFac.LoadFileLevel(rootPath, domain.GO)
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
