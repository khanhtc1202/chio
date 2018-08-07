package domain

type Program struct {
	RootPath string
}

type Module struct {
	SourceFiles []*SourceFile
	Language    LanguageType
}

type SourceFile struct {
	Path string
}

type Modules []*Module

func EmptyModuleList() *Modules {
	return &Modules{}
}
