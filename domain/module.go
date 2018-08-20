package domain

type ModuleProperties interface {
	GetFanInDepend() int
	GetFanOutDepend() int
	GetAbstractMember() int
	GetConcreteMember() int
}

// TODO implement for golang
type GolangModule struct {
	ModuleProperties
}

// TODO implement for python
type PythonModule struct {
	ModuleProperties
}
