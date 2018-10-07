package entity

import (
	"errors"
	"fmt"
)

type Modules map[string]*Module

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

func (m *Modules) LoadFileOfModule() {
	// TODO load files for each module
}
