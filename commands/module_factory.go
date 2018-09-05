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

func (m *ModuleFactory) FileAsModule(
	rootPath string,
	language domain.LanguageType,
) (*domain.Modules, error) {
	modules := domain.EmptyModuleList()
	module, err := m.loadFilesInDir(rootPath, language)
	if err != nil {
		return nil, err
	}
	modules.Add(module)
	return modules, nil
}

func (m *ModuleFactory) loadFilesInDir(
	dirPath string,
	language domain.LanguageType,
) (*domain.Module, error) {
	module := domain.NewModule(language)
	filepath.Walk(dirPath, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(language.Extension(), f.Name())
			if err == nil && r {
				module.AddSourceFile(domain.NewSourceFile(path))
			}
		}
		return nil
	})
	return module, nil
}
