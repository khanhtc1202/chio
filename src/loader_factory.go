package src

type Loader interface {
	CountConcreteMembers(files SourceFiles) (int, error)
	CountAbstractMembers(files SourceFiles) (int, error)
	ReferenceToPaths(files SourceFiles) ([]string, error)
}

type GoFileLoader struct {
}

func (g *GoFileLoader) CountConcreteMembers(files SourceFiles) (int, error) {
	// TODO implement
	return 0, nil
}

func (g *GoFileLoader) CountAbstractMembers(files SourceFiles) (int, error) {
	// TODO implement
	return 0, nil
}

func (g *GoFileLoader) ReferenceToPaths(files SourceFiles) ([]string, error) {
	// TODO implement
	return nil, nil
}

func LoaderFactory(lang LanguageType) Loader {
	switch lang {
	case GO:
		return &GoFileLoader{}
	default:
		return &GoFileLoader{}
	}
}
