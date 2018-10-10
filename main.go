package main

import (
	"os"
)

func main() {
	println("Hello")
	pwd, _ := os.Getwd()

	moduleFact := NewModuleFactory()

	modules, err := moduleFact.DirectoryAsModule(pwd, GO)
	if err != nil {
		panic("Error on load modules")
	}

	for _, module := range modules {
		loader := NewLoader(module.Language)
	}

	//println(pwd)
	//for _, module := range modules {
	//	files := module.GetSourceFilesPath()
	//	fmt.Println(files)
	//}
}
