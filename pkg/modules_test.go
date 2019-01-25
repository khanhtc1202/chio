package pkg_test

import (
	"testing"

	"github.com/khanhtc1202/chio/pkg"
	"github.com/stretchr/testify/assert"
)

func TestModules_Add_FailByAddEmptyModule(t *testing.T) {
	modules := pkg.NewModules()
	module := pkg.NewModule("/")

	err := modules.Add(module)
	assert.NotNil(t, err)
}

func TestModules_Load(t *testing.T) {
	modules := fakeLoadableModules()
	modules.Load(&MockLoaderRefToB{})

	moduleA := modules.GetModuleByPath("/src/a")
	assert.Equal(t, 2, moduleA.AbstractMember)
	assert.Equal(t, 2, moduleA.ConcreteMember)
	assert.Equal(t, 1, moduleA.FanOutDep)
	assert.Equal(t, 0, moduleA.FanInDep)

	moduleB := modules.GetModuleByPath("/src/b")
	assert.Equal(t, 2, moduleB.AbstractMember)
	assert.Equal(t, 2, moduleB.ConcreteMember)
	assert.Equal(t, 1, moduleB.FanOutDep) // +1 ref from A & +1 by impl of mock when load B
	assert.Equal(t, 2, moduleB.FanInDep)  // +1 by impl of mock when load B
}

func fakeLoadableModules() pkg.Modules {
	modules := pkg.NewModules()
	moduleB := pkg.NewModule("/src/b")
	moduleA := pkg.NewModule("/src/a")

	modules[moduleB.RootPath] = moduleB
	modules[moduleA.RootPath] = moduleA
	return modules
}

// ref to module B loader
type MockLoaderRefToB struct {
	pkg.Loader
}

func (m *MockLoaderRefToB) CountConcreteMembers(*pkg.Module) (int, error) {
	return 2, nil
}
func (m *MockLoaderRefToB) CountAbstractMembers(*pkg.Module) (int, error) {
	return 2, nil
}
func (m *MockLoaderRefToB) ReferenceToPaths(*pkg.Module) ([]string, error) {
	return []string{"/src/b"}, nil
}
