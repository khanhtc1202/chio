package loaders

import (
	"testing"

	"github.com/khanhtc1202/chio/pkg"
	"github.com/stretchr/testify/assert"
)

func TestGoFileLoader_CountConcreteMembers_Go(t *testing.T) {
	fileLoader := &GoFileLoader{}
	target, err := fileLoader.CountConcreteMembers(targetModule())

	assert.Nil(t, err)
	assert.Equal(t, 1, target)
}

func TestGoFileLoader_CountAbstractMembers_Go(t *testing.T) {
	fileLoader := &GoFileLoader{}
	target, err := fileLoader.CountAbstractMembers(targetModule())

	assert.Nil(t, err)
	assert.Equal(t, 0, target)
}

func TestGoFileLoader_ReferenceToPaths_Go(t *testing.T) {
	fileLoader := &GoFileLoader{}
	target, err := fileLoader.ReferenceToPaths(targetModule())

	assert.Nil(t, err)
	assert.Equal(t, "regexp", target[0])
}

func TestGoFileLoader_ReferenceToPaths_HasAliasImport(t *testing.T) {
	fileLoader := &GoFileLoader{}
	target, err := fileLoader.ReferenceToPaths(hasAliasImportModule())

	assert.Nil(t, err)
	assert.Equal(t, "fmt", target[0])
}

func targetModule() *pkg.Module {
	return &pkg.Module{
		SourceFiles: []*pkg.SourceFile{
			{
				Path: "./go_loader.go",
			},
		},
	}
}

func hasAliasImportModule() *pkg.Module {
	return &pkg.Module{
		SourceFiles: []*pkg.SourceFile{
			{
				Path: "../../sample/sample.go",
			},
		},
	}
}
