package entity

type Modules map[string]*Module

func (m *Modules) Add(rootPath string, module *Module) {
	(*m)[rootPath] = module
}

func (m *Modules) LoadFileOfModule() {
	// TODO load file for each module
}
