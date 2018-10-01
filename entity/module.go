package entity

type Module struct {
	ModuleProperties
	SourceFiles []*SourceFile
	Language    LanguageType
}

func NewModule(language LanguageType) *Module {
	return &Module{
		Language:         language,
		SourceFiles:      []*SourceFile{},
		ModuleProperties: FactoryModuleProperties(language),
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

// modules type
type Modules map[string]*Module

func EmptyModuleList() *Modules {
	return &Modules{}
}

func (m *Modules) Add(rootPath string, module *Module) {
	(*m)[rootPath] = module
}
