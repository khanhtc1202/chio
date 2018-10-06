package main

import (
	"os"

	"fmt"

	"github.com/khanhtc1202/chio/commands"
	"github.com/khanhtc1202/chio/entity"
)

func main() {
	println("Hello")
	pwd, _ := os.Getwd()

	moduleFact := commands.NewModuleFactory()

	modules, err := moduleFact.DirectoryAsModule(pwd, entity.GO)
	if err != nil {
		panic("Error on load modules")
	}

	println(pwd)
	for _, module := range modules {
		files := module.GetSourceFilesPath()
		fmt.Println(files)
	}
}
