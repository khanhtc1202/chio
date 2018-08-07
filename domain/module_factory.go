package domain

type ModuleFactory struct {
}

func (m *ModuleFactory) LoadFirstDirLevel(rootPath string) *Modules {
	modules := EmptyModuleList()
	// TODO impl for first dir level load strategy
	return modules
}

func (m *ModuleFactory) LoadFileLevel(rootPath string) *Modules {
	modules := EmptyModuleList()
	// TODO impl for file level load strategy
	return modules
}

func (m *ModuleFactory) LoadRecursionDirLevel(rootPath string) *Modules {
	modules := EmptyModuleList()
	// TODO impl for first dir level load strategy
	return modules
}
