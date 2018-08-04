package domain

type Modules []*Module

func EmptyModuleList() *Modules {
	return &Modules{}
}

func (m *Modules) Add(module *Module) {
	if module == nil {
		return
	}
	*m = append(*m, module)
}
