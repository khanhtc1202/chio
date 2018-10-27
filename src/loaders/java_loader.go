package loaders

import "github.com/khanhtc1202/chio/src"

type JavaFileLoader struct {
}

func (j *JavaFileLoader) CountConcreteMembers(module *src.Module) (int, error) {
	// TODO
	return 0, nil
}

func (j *JavaFileLoader) CountAbstractMembers(module *src.Module) (int, error) {
	// TODO
	return 0, nil
}

func (g *JavaFileLoader) ReferenceToPaths(module *src.Module) ([]string, error) {
	// TODO
	return []string{}, nil
}
