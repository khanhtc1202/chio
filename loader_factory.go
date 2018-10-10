package main

type Loader interface {
}

type GolangLoader struct {
	Loader
}

type JavaLoader struct {
	Loader
}

func NewLoader(lang LanguageType) Loader {
	switch lang {
	case GO:
		return &GolangLoader{}
	case JAVA:
		return &JavaLoader{}
	default:
		return &GolangLoader{}
	}
}
