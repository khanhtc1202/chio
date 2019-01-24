package src

import (
	"io/ioutil"
	"path/filepath"
)

type SourceFiles []*SourceFile

func EmptySourceFiles() SourceFiles {
	return []*SourceFile{}
}

type SourceFile struct {
	Path string
}

func NewSourceFile(path string) *SourceFile {
	absPath, _ := filepath.Abs(path)
	return &SourceFile{
		Path: absPath,
	}
}

func (s *SourceFile) Name() string {
	_, file := filepath.Split(s.Path)
	return file
}

func (s *SourceFile) DirPath() string {
	absPath, _ := filepath.Abs(s.Path)
	dir, _ := filepath.Split(absPath)
	return dir
}

func (s *SourceFile) Content() (string, error) {
	b, err := ioutil.ReadFile(s.Path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
