package domain

type Module interface {
	CountFanInDepend() int
	CountFanOutDepend() int
	CountAbstractMember() int
	CountConcreteMember() int
}

// TODO implement count method
type DirLevelModule struct {
	Module
}

// TODO implement count method
type FileLevelModule struct {
	Module
}

// TODO implement count method
type RecursionDirLevelModule struct {
	Module
}
