package src

import (
	"errors"
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
		if modules.GetModuleByPath(file.DirPath()) == nil {
			module := NewModule(file.DirPath())
			err = module.AddSourceFile(file)
			err = modules.Add(module)
		} else {
			err = modules.GetModuleByPath(file.DirPath()).AddSourceFile(file)
		}
	}
	return modules, err
}

func (m *ModuleFactory) loadModulesFiles(
	dirPath string,
	language LanguageType,
) (SourceFiles, error) {
	var files SourceFiles
	filepath.Walk(dirPath, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !f.IsDir() {
			r, err := regexp.MatchString(language.Extension(), f.Name())
			if err == nil && r {
				files = append(files, NewSourceFile(path))
			}
		}
		return nil
	})
	if files == nil {
		return nil, errors.New("try to load not existed or empty dir")
	}
	return files, nil
}
