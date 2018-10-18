package src

import (
	"errors"
	"fmt"
)

type Modules map[string]*Module

func NewModules() Modules {
	return Modules{}
}

func (m *Modules) Add(module *Module) error {
	if module.RootPath == "" || len(module.SourceFiles) == 0 {
		return errors.New(fmt.Sprintf("Error: Add empty module to modules list"))
	}

	(*m)[module.RootPath] = module
	return nil
}

func (m *Modules) GetModuleByPath(path string) *Module {
	return (*m)[path]
}

func (m *Modules) Load(loader Loader) error {
	for _, module := range *m {
		module.ConcreteMember, _ = loader.CountConcreteMembers(module)
		module.AbstractMember, _ = loader.CountAbstractMembers(module)
		refPaths, _ := loader.ReferenceToPaths(module)
		for _, path := range refPaths {
			module.FanInDep += 1
			if refModule := m.GetModuleByPath(path); refModule != nil {
				refModule.FanOutDep += 1
			}
		}
	}
	return nil
}
