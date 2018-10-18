package src_test

import (
	"testing"

	"github.com/khanhtc1202/chio/src"
	"github.com/stretchr/testify/assert"
)

func TestModules_Add_FailByAddEmptyModule(t *testing.T) {
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
	assert.Equal(t, 2, moduleA.FanInDep)
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
	moduleA := src.NewModule("/src/a", &MockLoaderRefToBAndExternalModule{})

	modules[moduleB.RootPath] = moduleB
	modules[moduleA.RootPath] = moduleA
	return modules
}

// ref to module B loader
type MockLoaderRefToBAndExternalModule struct {
	src.Loader
}

func (m *MockLoaderRefToBAndExternalModule) CountConcreteMembers() (int, error) {
	return 2, nil
}
func (m *MockLoaderRefToBAndExternalModule) CountAbstractMembers() (int, error) {
	return 2, nil
}
func (m *MockLoaderRefToBAndExternalModule) ReferenceToModules() ([]string, error) {
	return []string{"/src/b", "/src/external"}, nil
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
