package src_test

import (
	"testing"

	"github.com/khanhtc1202/chio/src"
	"github.com/stretchr/testify/assert"
)

func TestModules_Add_AddNotEmptyModule(t *testing.T) {
	modules := src.NewModules()
	module := src.NewModule(nil)
	module.RootPath = "/"

	err := modules.Add(module)
	assert.Nil(t, err)
}

func TestModules_Add_AddEmptyModule(t *testing.T) {
	modules := src.NewModules()
	module := src.NewModule(nil)

	err := modules.Add(module)
	assert.NotNil(t, err)
}

func TestModules_Load(t *testing.T) {
	modules := loadableModules
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

var loadableModules = fakeLoadableModules()

func fakeLoadableModules() src.Modules {
	modules := src.NewModules()
	moduleB := fakeLoadedModule("/src/b", &MockLoaderRefNil{})
	moduleA := fakeLoadedModule("/src/a", &MockLoaderRefToB{})

	modules.Add(moduleA)
	modules.Add(moduleB)
	return modules
}

func fakeLoadedModule(path string, loader src.Loader) *src.Module {
	return &src.Module{
		Loader:   loader,
		RootPath: path,
	}
}

func fakeRefToModule(mm *src.Module) []*src.Module {
	return []*src.Module{
		mm,
	}
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
func (m *MockLoaderRefToB) ReferenceToModules() ([]*src.Module, error) {
	//moduleB := fakeLoadedModule("/src/b", &MockLoaderRefNil{})
	moduleB := loadableModules.GetModuleByPath("/src/b")
	return fakeRefToModule(moduleB), nil
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
func (m *MockLoaderRefNil) ReferenceToModules() ([]*src.Module, error) {
	return nil, nil
}
