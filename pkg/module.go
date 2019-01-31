package pkg

import (
	"errors"
	"fmt"
	"math"
	"path/filepath"
	"strings"
)

type Module struct {
	RootPath       string
	SourceFiles    SourceFiles
	FanInDep       int
	FanOutDep      int
	AbstractMember int
	ConcreteMember int
}

func NewModule(rootPath string) *Module {
	absPath, _ := filepath.Abs(rootPath)
	return &Module{
<<<<<<< HEAD
		RootPath:    absPath + "/", // module dir path end with '/'
=======
		RootPath:    absPath,
>>>>>>> 7cae6a0c7f1057cbc163dd35969511074d6f8c1e
		SourceFiles: EmptySourceFiles(),
	}
}

func (m *Module) AddSourceFile(file *SourceFile) error {
	if strings.Contains(file.Path, m.RootPath) {
		m.SourceFiles = append(m.SourceFiles, file)
		return nil
	}

	return errors.New(fmt.Sprintf("add not contain file path to module: %s\n", m.RootPath))
}

func (m *Module) AddSourceFiles(files SourceFiles) error {
	for _, file := range files {
		err := m.AddSourceFile(file)
		if err != nil {
			return err
		}
	}
	return nil
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
