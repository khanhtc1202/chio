package loaders

import (
	"regexp"
	"strings"

	"github.com/khanhtc1202/chio/src"
)

var (
	concreteMemberDeclareRegex = regexp.MustCompile(`(?:type)(\s)*\w+(\s)*(?:struct)(\s)*(?:\{)`)
	abstractMemberDeclareRegex = regexp.MustCompile(`(?:type)(\s)*\w+(\s)*(?:interface)(\s)*(?:\{)`)
	importStatementRegex       = regexp.MustCompile(`(\".+\"\n)|(\n(\s*)\".+\"\n)`)
	normalizeRegex             = regexp.MustCompile(`\"(.+?)\"`)
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

		matches := concreteMemberDeclareRegex.FindAllStringIndex(content, -1)
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

		matches := abstractMemberDeclareRegex.FindAllStringIndex(content, -1)
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

		matches := importStatementRegex.FindAllStringSubmatch(content, -1)
		for i := range matches {
			refPaths = append(refPaths, normalizeRegex.FindStringSubmatch(matches[i][0])[1])
		}
	}
	return refPaths, nil
}
