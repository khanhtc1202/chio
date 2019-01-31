package pkg

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

type NDepthDirectoryAsModule struct {
	moduleFactory
}

func NewNDepthDirectoryAsModule(path string) *NDepthDirectoryAsModule {
	return &NDepthDirectoryAsModule{
		moduleFactory{modulePath: path},
	}
}

func (m *NDepthDirectoryAsModule) Load(
	language LanguageType,
) (Modules, error) {
	modules := NewModules()
	files, err := m.loadModulesFiles(language)
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

type OneDepthDirectoryAsModule struct {
	moduleFactory
}

func NewOneDepthDirectoryAsModule(path string) *OneDepthDirectoryAsModule {
	return &OneDepthDirectoryAsModule{
		moduleFactory{modulePath: path},
	}
}

func (m *OneDepthDirectoryAsModule) Load(
	language LanguageType,
) (Modules, error) {
	modules := NewModules()

	ioFiles, err := ioutil.ReadDir(m.modulePath)
	if err != nil {
		return nil, err
	}

	for _, fileInfo := range ioFiles {
		if fileInfo.IsDir() {
			subModulePath := filepath.Join(m.modulePath, fileInfo.Name())
			fac := &moduleFactory{modulePath: subModulePath}
			files, err := fac.loadModulesFiles(language)
			if err != nil {
				return nil, err
			}
			module := NewModule(subModulePath)
			err = module.AddSourceFiles(files)
			if err != nil {
				return nil, err
			}
			err = modules.Add(module)
		}
	}

	return modules, nil
}

type moduleFactory struct {
	modulePath string
}

func (m *moduleFactory) loadModulesFiles(
	language LanguageType,
) (SourceFiles, error) {
	var files SourceFiles
	filepath.Walk(m.modulePath, func(path string, f os.FileInfo, err error) error {
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
