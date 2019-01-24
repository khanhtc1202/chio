package src

import (
	"errors"
	"strings"
)

type Modules map[string]*Module

func NewModules() Modules {
	return Modules{}
}

func (m *Modules) Add(module *Module) error {
	if module.RootPath == "" || len(module.SourceFiles) == 0 {
		return errors.New("add empty module to modules list")
	}

	(*m)[module.RootPath] = module
	return nil
}

func (m *Modules) GetModuleByPath(path string) *Module {
	for mPath, module := range *m {
		if strings.Contains(mPath, path) {
			return module
		}
	}
	return nil
}

func (m *Modules) Load(loader Loader) (err error) {
	for _, module := range *m {
		module.ConcreteMember, err = loader.CountConcreteMembers(module)
		module.AbstractMember, err = loader.CountAbstractMembers(module)
		refPaths, err := loader.ReferenceToPaths(module)
		if err != nil {
			return err
		}
		for _, path := range refPaths {
			module.FanOutDep += 1
			if refModule := m.GetModuleByPath(path); refModule != nil {
				refModule.FanInDep += 1
			}
		}
	}
	return nil
}
