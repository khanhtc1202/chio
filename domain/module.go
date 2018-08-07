package domain

type ModuleProperties interface {
	CountFanInDepend() int
	CountFanOutDepend() int
	CountAbstractMember() int
	CountConcreteMember() int
}

// TODO implement for golang
type GolangModule struct {
	ModuleProperties
}

// TODO implement for python
type PythonModule struct {
	ModuleProperties
}
