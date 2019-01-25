package pkg

type Loader interface {
	CountConcreteMembers(*Module) (int, error)
	CountAbstractMembers(*Module) (int, error)
	ReferenceToPaths(*Module) ([]string, error)
}
