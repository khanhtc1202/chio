package src_test

import (
	"testing"

	"github.com/khanhtc1202/chio/src"
	"github.com/stretchr/testify/assert"
)

func TestModules_Add_FailByAddEmptyModule(t *testing.T) {
	modules := src.NewModules()
	module := src.NewModule("/")

	err := modules.Add(module)
	assert.NotNil(t, err)
}

func TestModules_GetModuleByPath_GetSubPath(t *testing.T) {
}

func TestModules_Load(t *testing.T) {
	modules := fakeLoadableModules()
	modules.Load(&MockLoaderRefToB{})

	moduleA := modules.GetModuleByPath("/src/a")
	assert.Equal(t, 2, moduleA.AbstractMember)
	assert.Equal(t, 2, moduleA.ConcreteMember)
	assert.Equal(t, 1, moduleA.FanInDep)
	assert.Equal(t, 0, moduleA.FanOutDep)

	moduleB := modules.GetModuleByPath("/src/b")
	assert.Equal(t, 2, moduleB.AbstractMember)
	assert.Equal(t, 2, moduleB.ConcreteMember)
	assert.Equal(t, 1, moduleB.FanInDep)  // +1 by impl of mock when load B
	assert.Equal(t, 2, moduleB.FanOutDep) // +1 ref from A & +1 by impl of mock when load B
}

func fakeLoadableModules() src.Modules {
	modules := src.NewModules()
	moduleB := src.NewModule("/src/b")
	moduleA := src.NewModule("/src/a")

	modules[moduleB.RootPath] = moduleB
	modules[moduleA.RootPath] = moduleA
	return modules
}

// ref to module B loader
type MockLoaderRefToB struct {
	src.Loader
}

func (m *MockLoaderRefToB) CountConcreteMembers(*src.Module) (int, error) {
	return 2, nil
}
func (m *MockLoaderRefToB) CountAbstractMembers(*src.Module) (int, error) {
	return 2, nil
}
func (m *MockLoaderRefToB) ReferenceToPaths(*src.Module) ([]string, error) {
	return []string{"/src/b"}, nil
}
