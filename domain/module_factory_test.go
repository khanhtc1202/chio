package domain_test

import (
	"testing"

	"fmt"

	"github.com/khanhtc1202/chio/domain"
)

func TestModuleFactory_LoadFileLevel(t *testing.T) {
	moduleFac := domain.NewModuleFactory()
	modules, err := moduleFac.LoadFileLevel("../../", domain.GO)
	if err != nil {
		t.Fatal("Error on load file level test:", err)
	}
	if len(*modules) == 0 {
		t.Fatal("Count false in load module file level")
	}
}

func TestModuleFactory_LoadFirstDirLevel(t *testing.T) {
	moduleFac := domain.NewModuleFactory()
	modules, err := moduleFac.LoadFirstDirLevel("/Users/khanh.tran/workspace/go/src/github.com/khanhtc1202/chio", domain.GO)
	if err != nil {
		t.Fatal("Error on load file by first dir level test:", err)
	}
	if len(*modules) == 0 {
		t.Fatal("Count false in load module first dir level")
	}
	for _, module := range *modules {
		files := module.GetSourceFilesPath()
		fmt.Println(files)
	}
}

func TestModuleFactory_LoadRecursionDirLevel(t *testing.T) {
	moduleFac := domain.NewModuleFactory()
	modules, err := moduleFac.LoadRecursionDirLevel("../../", domain.GO)
	if err != nil {
		t.Fatal("Error on load file recursion dir level test:", err)
	}
	if len(*modules) == 0 {
		t.Fatal("Count false in load module recursion dir level")
	}
}
