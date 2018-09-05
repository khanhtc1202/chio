package domain

type Program struct {
	RootPath string
}

type Modules []*Module

func EmptyModuleList() *Modules {
	return &Modules{}
}

func (m *Modules) Add(module *Module) {
	*m = append(*m, module)
}
