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
	if module.RootPath != "" {
		(*m)[module.RootPath] = module
		return nil
	}

	return errors.New(fmt.Sprintf("Error: Add empty module to modules list"))
}

func (m *Modules) GetModuleByPath(path string) *Module {
	return (*m)[path]
}

func (m *Modules) Load() error {
	for _, module := range *m {
		module.ConcreteMember, _ = module.CountConcreteMembers()
		module.AbstractMember, _ = module.CountAbstractMembers()
		refToModules, _ := module.ReferenceToModules()
		for _, refModule := range refToModules {
			module.FanInDep += 1
			refModule.FanOutDep += 1
		}
	}
	return nil
}
