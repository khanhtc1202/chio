package main

import (
	"fmt"

	"flag"

	"github.com/khanhtc1202/chio/src"
)

type CommandParams struct {
	Path     string
	Language string
}

func parseParams() *CommandParams {
	var modulePath string
	flag.StringVar(&modulePath, "p", ".", "path to module")
	language := flag.String("l", "go", "language(s): go")

	flag.Parse()

	return &CommandParams{
		Path:     modulePath,
		Language: *language,
	}
}

func main() {
	cmdParams := parseParams()

	modulePath := cmdParams.Path
	language := src.ValueOfLanguage(cmdParams.Language)

	// load dir level
	// TODO move to command params
	moduleFact := src.NewModuleFactory()
	modules, err := moduleFact.DirectoryAsModule(modulePath, language)
	if err != nil {
		panic("Error on load modules")
	}

	// loader object by language
	loader := src.LoaderFactory(language)

	// load modules
	modules.Load(loader)

	// print output
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
