package domain

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

func (m *ModuleFactory) LoadFirstDirLevel(rootPath string, language LanguageType) (*Modules, error) {
	modules := EmptyModuleList()
	// TODO impl for first dir level load strategy
	module := NewModule(language)
	modules.Add(module)
	filepath.Walk(rootPath, func(path string, f os.FileInfo, _ error) error {
		if f.IsDir() {
			module = NewModule(language)
			modules.Add(module)
		} else {
			r, err := regexp.MatchString(language.Extension(), f.Name())
			if err == nil && r {
				module.AddSourceFile(NewSourceFile(f.Name()))
			}
		}
		return nil
	})
	return modules, nil
}

func (m *ModuleFactory) LoadFileLevel(rootPath string, language LanguageType) (*Modules, error) {
	modules := EmptyModuleList()
	// TODO impl for file level load strategy
	return modules, nil
}

func (m *ModuleFactory) LoadRecursionDirLevel(rootPath string, language LanguageType) (*Modules, error) {
	modules := EmptyModuleList()
	// TODO impl for first dir level load strategy
	return modules, nil
}
