package main

import (
	"os"

	"github.com/khanhtc1202/chio/src"
)

func main() {
	pwd, _ := os.Getwd()

	moduleFact := src.NewModuleFactory()

	_, err := moduleFact.DirectoryAsModule(pwd, src.GO)
	if err != nil {
		panic("Error on load modules")
	}

	//for _, module := range modules {
	//	loader := NewLoader(module.Language)
	//}

	//println(pwd)
	//for _, module := range modules {
	//	files := module.GetSourceFilesPath()
	//	fmt.Println(files)
	//}
}
