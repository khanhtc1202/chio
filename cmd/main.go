package main

import (
	"fmt"
	"path/filepath"

	"github.com/fatih/color"

	"flag"

	"os"
	"strings"

	"github.com/khanhtc1202/chio/pkg"
	"github.com/khanhtc1202/chio/pkg/loaders"
	"github.com/olekukonko/tablewriter"
)

type CommandParams struct {
	Path     string
	Language string
	Depth    string
}

func parseParams() *CommandParams {
	var modulePath string
	var language string
	var depth string

	flag.StringVar(&modulePath, "p", ".", "path to module")
	flag.StringVar(&language, "l", "go", "language(s): go")
	flag.StringVar(&depth, "d", "n", "dir as module, default n-depth (n)")

	flag.Parse()

	return &CommandParams{
		Path:     modulePath,
		Language: language,
		Depth:    depth,
	}
}

func colorfulMetric(data float64) string {
	if data < 0.3 {
		return color.GreenString("%.3f", data)
	}
	if data < 0.6 {
		return color.YellowString("%.3f", data)
	}
	return color.RedString("%.3f", data)
}

func print(srcPath string, modules pkg.Modules) {
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
			colorfulMetric(module.Distance()),
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

	srcPath, _ := filepath.Abs(cmdParams.Path)
	language := pkg.ValueOfLanguage(cmdParams.Language)
	moduleFact := pkg.ModuleFactoryBySign(cmdParams.Depth, srcPath)

	modules, err := moduleFact.Load(language)
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
