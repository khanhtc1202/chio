package main

import (
	"fmt"

	"flag"

	"os"
	"strings"

	"github.com/khanhtc1202/chio/src"
	"github.com/khanhtc1202/chio/src/loaders"
	"github.com/olekukonko/tablewriter"
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

func print(srcPath string, modules src.Modules) {
	var data [][]string

	for _, module := range modules {
		row := []string{
			strings.Replace(module.RootPath, srcPath, "", -1),
			fmt.Sprintf("%d", len(module.SourceFiles)),
			fmt.Sprintf("%d", module.ConcreteMember),
			fmt.Sprintf("%d", module.AbstractMember),
			fmt.Sprintf("%d", module.FanInDep),
			fmt.Sprintf("%d", module.FanOutDep),
			fmt.Sprintf("%.3f", module.Abstractness()),
			fmt.Sprintf("%.3f", module.Instability()),
			fmt.Sprintf("%.3f", module.Distance()),
		}

		data = append(data, row)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Module Path", "Files", "Concrete", "Abstract", "FanIn", "FanOut", "Abstractness", "Instability", "Distance"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}

func main() {
	cmdParams := parseParams()

	srcPath := cmdParams.Path
	language := src.ValueOfLanguage(cmdParams.Language)

	// load dir level
	// TODO move to command params
	moduleFact := src.NewModuleFactory()
	modules, err := moduleFact.DirectoryAsModule(srcPath, language)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}

	// loader object by language
	loader := loaders.LoaderFactory(language)

	// load modules
	modules.Load(loader)

	// print output
	print(srcPath, modules)
}
