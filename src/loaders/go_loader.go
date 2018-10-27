package loaders

import (
	"regexp"
	"strings"

	"github.com/khanhtc1202/chio/src"
)

type GoFileLoader struct {
}

func (g *GoFileLoader) CountConcreteMembers(module *src.Module) (int, error) {
	totalConcreteMembers := 0
	for _, file := range module.SourceFiles {
		if strings.Contains(file.Name(), "test") {
			continue
		}

		content, err := file.Content()
		if err != nil {
			return 0, err
		}

		concreteRex := regexp.MustCompile(`(?:type)(\s)*\w+(\s)*(?:struct)(\s)*(?:\{)`)
		matches := concreteRex.FindAllStringIndex(content, -1)
		totalConcreteMembers += len(matches)
	}
	return totalConcreteMembers, nil
}

func (g *GoFileLoader) CountAbstractMembers(module *src.Module) (int, error) {
	totalAbstractMembers := 0
	for _, file := range module.SourceFiles {
		if strings.Contains(file.Name(), "test") {
			continue
		}

		content, err := file.Content()
		if err != nil {
			return 0, err
		}

		abstractRex := regexp.MustCompile(`(?:type)(\s)*\w+(\s)*(?:interface)(\s)*(?:\{)`)
		matches := abstractRex.FindAllStringIndex(content, -1)
		totalAbstractMembers += len(matches)
	}
	return totalAbstractMembers, nil
}

func (g *GoFileLoader) ReferenceToPaths(module *src.Module) ([]string, error) {
	var refPaths []string
	for _, file := range module.SourceFiles {
		if strings.Contains(file.Name(), "test") {
			continue
		}

		content, err := file.Content()
		if err != nil {
			return refPaths, err
		}

		importRegex := regexp.MustCompile(`(\".+\"\n)|(\n(\s*)\".+\"\n)`)
		matches := importRegex.FindAllStringSubmatch(content, -1)
		for i := range matches {
			normalizeRegex := regexp.MustCompile(`\"(.+?)\"`)
			refPaths = append(refPaths, normalizeRegex.FindStringSubmatch(matches[i][0])[1])
		}
	}
	return refPaths, nil
}
