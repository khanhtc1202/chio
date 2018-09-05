package domain

type ModuleProperties interface {
	GetFanInDepend() int
	GetFanOutDepend() int
	GetAbstractMember() int
	GetConcreteMember() int
}

type Module struct {
	ModuleProperties
	SourceFiles []*SourceFile
	Language    LanguageType
}

func NewModule(language LanguageType) *Module {
	return &Module{
		SourceFiles: []*SourceFile{},
		Language:    language,
	}
}

func (m *Module) AddSourceFile(file *SourceFile) {
	m.SourceFiles = append(m.SourceFiles, file)
}

func (m *Module) GetSourceFilesPath() []string {
	var filesPath []string
	for _, path := range m.SourceFiles {
		filesPath = append(filesPath, path.Path)
	}
	return filesPath
}
