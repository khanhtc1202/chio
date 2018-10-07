package commands

import (
	"os"
	"path/filepath"
	"regexp"

	"github.com/khanhtc1202/chio/entity"
)

type ModuleFactory struct {
}

func NewModuleFactory() *ModuleFactory {
	return &ModuleFactory{}
}

func (m *ModuleFactory) DirectoryAsModule(
	rootPath string,
	language entity.LanguageType,
) (entity.Modules, error) {
	var modules entity.Modules
	files, err := m.loadModulesFiles(rootPath, language)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if modules.GetModuleByPath(file.GetDirPath()) == nil {
			module := entity.NewModule(language)
			module.AddSourceFile(file)
			modules.Add(module)
		} else {
			modules.GetModuleByPath(file.GetDirPath()).AddSourceFile(file)
		}
	}
	return modules, nil
}

func (m *ModuleFactory) loadModulesFiles(
	dirPath string,
	language entity.LanguageType,
) ([]*entity.SourceFile, error) {
	var files []*entity.SourceFile
	filepath.Walk(dirPath, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(language.Extension(), f.Name())
			if err == nil && r {
				files = append(files, entity.NewSourceFile(path))
			}
		}
		return nil
	})
	return files, nil
}
