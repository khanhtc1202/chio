package commands

import (
	"os"
	"path/filepath"
	"regexp"

	"github.com/khanhtc1202/chio/domain"
)

type ModuleFactory struct {
}

func NewModuleFactory() *ModuleFactory {
	return &ModuleFactory{}
}

func (m *ModuleFactory) DirectoryAsModule(
	rootPath string,
	language domain.LanguageType,
) (*domain.Modules, error) {
	modules := domain.EmptyModuleList()
	files, err := m.loadFilesInDir(rootPath, language)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if (*modules)[file.GetDirPath()] == nil {
			module := domain.NewModule(language)
			module.AddSourceFile(file)
			modules.Add(file.GetDirPath(), module)
		} else {
			(*modules)[file.GetDirPath()].AddSourceFile(file)
		}
	}
	return modules, nil
}

func (m *ModuleFactory) loadFilesInDir(
	dirPath string,
	language domain.LanguageType,
) ([]*domain.SourceFile, error) {
	var files []*domain.SourceFile
	filepath.Walk(dirPath, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(language.Extension(), f.Name())
			if err == nil && r {
				files = append(files, domain.NewSourceFile(path))
			}
		}
		return nil
	})
	return files, nil
}
