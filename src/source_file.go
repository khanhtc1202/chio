package src

import (
	"io/ioutil"
	"path"
)

type SourceFile struct {
	Path string
}

func NewSourceFile(path string) *SourceFile {
	return &SourceFile{
		Path: path,
	}
}

func (s *SourceFile) GetDirPath() string {
	dir, _ := path.Split(s.Path)
	return dir
}

func (s *SourceFile) Content() (string, error) {
	b, err := ioutil.ReadFile(s.Path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
