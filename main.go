package main

import (
	"fmt"

	"github.com/khanhtc1202/chio/src"
)

func main() {
	pwd := "/Users/khanh.tran/workspace/go/src/github.com/khanhtc1202/chio"

	moduleFact := src.NewModuleFactory()

	modules, err := moduleFact.DirectoryAsModule(pwd, src.GO)
	if err != nil {
		panic("Error on load modules")
	}

	loader := src.LoaderFactory(src.GO)

	modules.Load(loader)

	for _, module := range modules {
		fmt.Println("---------------")
		fmt.Printf("\nPath: %s", module.RootPath)
		fmt.Printf("\nConcrete/Abstract members: %d - %d", module.ConcreteMember, module.AbstractMember)
		fmt.Printf("\nFanIn/FanOut dependencies: %d - %d", module.FanInDep, module.FanOutDep)
		fmt.Printf("\nAbstractness: %v", module.Abstractness())
		fmt.Printf("\nInstability: %v", module.Instability())
		fmt.Printf("\nDistance: %v", module.Distance())
		fmt.Println("\n---------------")
	}
}
