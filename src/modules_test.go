package src_test

import (
	"testing"

	"github.com/khanhtc1202/chio/src"
	"github.com/stretchr/testify/assert"
)

func TestModules_Add_AddNotEmptyModule(t *testing.T) {
	modules := src.NewModules()
	module := src.NewModule("/", nil)

	err := modules.Add(module)
	assert.NotNil(t, err)
}

func TestModules_Load(t *testing.T) {
	modules := fakeLoadableModules()
	modules.Load()

	moduleA := modules.GetModuleByPath("/src/a")
	assert.Equal(t, 2, moduleA.AbstractMember)
	assert.Equal(t, 2, moduleA.ConcreteMember)
	assert.Equal(t, 1, moduleA.FanInDep)
	assert.Equal(t, 0, moduleA.FanOutDep)

	moduleB := modules.GetModuleByPath("/src/b")
	assert.Equal(t, 2, moduleB.AbstractMember)
	assert.Equal(t, 2, moduleB.ConcreteMember)
	assert.Equal(t, 0, moduleB.FanInDep)
	assert.Equal(t, 1, moduleB.FanOutDep)
}

func fakeLoadableModules() src.Modules {
	modules := src.NewModules()
	moduleB := src.NewModule("/src/b", &MockLoaderRefNil{})
	moduleA := src.NewModule("/src/a", &MockLoaderRefToB{})

	modules[moduleB.RootPath] = moduleB
	modules[moduleA.RootPath] = moduleA
	return modules
}

// ref to module B loader
type MockLoaderRefToB struct {
	src.Loader
}

func (m *MockLoaderRefToB) CountConcreteMembers() (int, error) {
	return 2, nil
}
func (m *MockLoaderRefToB) CountAbstractMembers() (int, error) {
	return 2, nil
}
func (m *MockLoaderRefToB) ReferenceToModules() ([]string, error) {
	return []string{"/src/b"}, nil
}

// ref to nil module loader
type MockLoaderRefNil struct {
	src.Loader
}

func (m *MockLoaderRefNil) CountConcreteMembers() (int, error) {
	return 2, nil
}
func (m *MockLoaderRefNil) CountAbstractMembers() (int, error) {
	return 2, nil
}
func (m *MockLoaderRefNil) ReferenceToModules() ([]string, error) {
	return []string{}, nil
}
