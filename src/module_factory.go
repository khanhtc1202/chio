package src

import (
	"os"
	"path/filepath"
	"regexp"
)

type ModuleFactory struct {
}

func NewModuleFactory() *ModuleFactory {
	return &ModuleFactory{}
}

func (m *ModuleFactory) DirectoryAsModule(
	rootPath string,
	language LanguageType,
) (Modules, error) {
	modules := NewModules()
	files, err := m.loadModulesFiles(rootPath, language)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if modules.GetModuleByPath(file.GetDirPath()) == nil {
			module := NewModule(file.GetDirPath(), LoaderFactory(language))
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
	language LanguageType,
) ([]*SourceFile, error) {
	var files []*SourceFile
	filepath.Walk(dirPath, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(language.Extension(), f.Name())
			if err == nil && r {
				files = append(files, NewSourceFile(path))
			}
		}
		return nil
	})
	return files, nil
}
