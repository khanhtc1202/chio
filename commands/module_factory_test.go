package commands

import (
	"fmt"
	"testing"

	"os"
	"path/filepath"

	"github.com/khanhtc1202/chio/entity"
)

func TestModuleFactory_DirectoryAsModuleFileLevel(t *testing.T) {
	rootPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	moduleFac := NewModuleFactory()
	modules, err := moduleFac.DirectoryAsModule(rootPath, entity.GO)
	if err != nil {
		t.Fatal("Error on load file level test:", err)
	}
	if len(modules) == 0 {
		t.Fatal("Count false in load module file level")
	}
	for _, module := range modules {
		files := module.GetSourceFilesPath()
		fmt.Println(files)
	}
}
