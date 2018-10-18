package src_test

import (
	"testing"

	"github.com/khanhtc1202/chio/src"
	"github.com/stretchr/testify/assert"
)

func TestGoFileLoader_CountConcreteMembers_Go(t *testing.T) {
	fileLoader := &src.GoFileLoader{}
	target, err := fileLoader.CountConcreteMembers(targetModule())

	assert.Nil(t, err)
	assert.Equal(t, 1, target)
}

func TestGoFileLoader_CountAbstractMembers_Go(t *testing.T) {
	fileLoader := &src.GoFileLoader{}
	target, err := fileLoader.CountAbstractMembers(targetModule())

	assert.Nil(t, err)
	assert.Equal(t, 1, target)
}

func TestGoFileLoader_ReferenceToPaths_Go(t *testing.T) {
	fileLoader := &src.GoFileLoader{}
	target, err := fileLoader.ReferenceToPaths(targetModule())

	assert.Nil(t, err)
	assert.Equal(t, "regexp", target[0])
}

func targetModule() *src.Module {
	return &src.Module{
		SourceFiles: []*src.SourceFile{
			{
				Path: "./loader_factory.go",
			},
		},
	}
}
