package src

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

type Module struct {
	Loader
	RootPath       string
	SourceFiles    []*SourceFile
	FanInDep       int
	FanOutDep      int
	AbstractMember int
	ConcreteMember int
}

func NewModule(rootPath string, loader Loader) *Module {
	return &Module{
		Loader:      loader,
		RootPath:    rootPath,
		SourceFiles: []*SourceFile{},
	}
}

func (m *Module) AddSourceFile(file *SourceFile) error {
	if strings.Contains(file.Path, m.RootPath) {
		m.SourceFiles = append(m.SourceFiles, file)
		return nil
	}

	return errors.New(fmt.Sprintf("Error: Add not contain file path to module: %s\n", m.RootPath))
}

func (m *Module) GetSourceFilesPath() []string {
	var filesPath []string
	for _, path := range m.SourceFiles {
		filesPath = append(filesPath, path.Path)
	}
	return filesPath
}

func (m *Module) Instability() float64 {
	fOut := m.FanOutDep
	fIn := m.FanInDep
	return float64(fOut) / float64(fOut+fIn)
}

func (m *Module) Abstractness() float64 {
	abs := m.AbstractMember
	con := m.ConcreteMember
	return float64(abs) / float64(abs+con)
}

func (m *Module) Distance() float64 {
	// cal instability
	fOut := m.FanOutDep
	fIn := m.FanInDep
	instability := float64(fOut) / float64(fOut+fIn)

	// cal abstractness
	abs := m.AbstractMember
	con := m.ConcreteMember
	abstractness := float64(abs) / float64(abs+con)

	return math.Abs(instability + abstractness - 1)
}
