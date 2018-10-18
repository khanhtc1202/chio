package src

type Loader interface {
	CountConcreteMembers() (int, error)
	CountAbstractMembers() (int, error)
	ReferenceToPaths() ([]string, error)
}

type GoFileLoader struct {
	Loader
}

func (g *GoFileLoader) CountConcreteMembers() (int, error) {
	// TODO implement
	return 0, nil
}

func (g *GoFileLoader) CountAbstractMembers() (int, error) {
	// TODO implement
	return 0, nil
}

func (g *GoFileLoader) ReferenceToPaths() ([]string, error) {
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
