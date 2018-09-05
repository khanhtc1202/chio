package domain

import "io/ioutil"

type Program struct {
	RootPath string
}

type Module struct {
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

type SourceFile struct {
	Path string
}

func NewSourceFile(path string) *SourceFile {
	return &SourceFile{
		Path: path,
	}
}

func (s *SourceFile) Content() string {
	b, err := ioutil.ReadFile(s.Path)
	if err != nil {
		// TODO return on err load content file
		return ""
	}
	return string(b)
}

type Modules []*Module

func EmptyModuleList() *Modules {
	return &Modules{}
}

func (m *Modules) Add(module *Module) {
	*m = append(*m, module)
}
