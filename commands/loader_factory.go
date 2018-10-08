package commands

import (
	"github.com/khanhtc1202/chio/entity"
)

type Loader interface {
}

type GolangLoader struct {
	Loader
}

type JavaLoader struct {
	Loader
}

func NewLoader(lang entity.LanguageType) Loader {
	switch lang {
	case entity.GO:
		return &GolangLoader{}
	case entity.JAVA:
		return &JavaLoader{}
	default:
		return &GolangLoader{}
	}
}
