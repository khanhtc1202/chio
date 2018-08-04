package domain

type Program interface {
	GetModuleList() Modules
}

type GolangProgram struct {
}

func (*GolangProgram) GetModuleList() Modules {
	// TODO
	return nil
}

type PythonProgram struct {
}

func (*PythonProgram) GetModuleList() Modules {
	// TODO
	return nil
}
