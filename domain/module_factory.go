package domain

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// TODO move to config
var ignoreDirs = []string{".git", ".idea"}

type ModuleFactory struct {
}

func NewModuleFactory() *ModuleFactory {
	return &ModuleFactory{}
}

func (m *ModuleFactory) LoadFirstDirLevel(rootPath string, language LanguageType) (*Modules, error) {
	modules := EmptyModuleList()
	baseDirs, err := m.buildBaseDirs(rootPath)
	if err != nil {
		return modules, err
	}

	for _, dir := range baseDirs {
		module, err := m.loadFilesInDir(dir, language)
		if err != nil {
			return modules, err
		}
		if len(module.SourceFiles) > 0 {
			modules.Add(module)
		}
	}
	return modules, nil
}

func (m *ModuleFactory) LoadFileLevel(rootPath string, language LanguageType) (*Modules, error) {
	modules := EmptyModuleList()
	module, err := m.loadFilesInDir(rootPath, language)
	if err != nil {
		return nil, err
	}
	modules.Add(module)
	return modules, nil
}

func (m *ModuleFactory) LoadRecursionDirLevel(rootPath string, language LanguageType) (*Modules, error) {
	modules := EmptyModuleList()
	// TODO impl for first dir level load strategy
	return modules, nil
}

func (m *ModuleFactory) loadFilesInDir(dirPath string, language LanguageType) (*Module, error) {
	module := NewModule(language)
	filepath.Walk(dirPath, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(language.Extension(), f.Name())
			if err == nil && r {
				module.AddSourceFile(NewSourceFile(path))
			}
		}
		return nil
	})
	return module, nil
}

func (m *ModuleFactory) buildBaseDirs(rootPath string) ([]string, error) {
	dirs, err := ioutil.ReadDir(rootPath)
	if err != nil {
		return nil, err
	}

	var baseDirs []string
	// TODO move to config
	ignoreDirsPattern := strings.Join(ignoreDirs, "")

	for _, dir := range dirs {
		r, err := regexp.MatchString(dir.Name(), ignoreDirsPattern)
		if err == nil && r {
			continue
		}
		baseDirs = append(baseDirs, rootPath+"/"+dir.Name())
	}
	return baseDirs, nil
}
