package loaders

import "github.com/khanhtc1202/chio/pkg"

type JavaFileLoader struct {
}

func (j *JavaFileLoader) CountConcreteMembers(module *pkg.Module) (int, error) {
	// TODO
	return 0, nil
}

func (j *JavaFileLoader) CountAbstractMembers(module *pkg.Module) (int, error) {
	// TODO
	return 0, nil
}

func (g *JavaFileLoader) ReferenceToPaths(module *pkg.Module) ([]string, error) {
	// TODO
	return []string{}, nil
}
