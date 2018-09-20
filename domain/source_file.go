package domain

import (
	"io/ioutil"
	"os"
)

type SourceFile struct {
	Path string
}

func NewSourceFile(path string) *SourceFile {
	return &SourceFile{
		Path: path,
	}
}

func (s *SourceFile) Content() (string, error) {
	b, err := ioutil.ReadFile(s.Path)
	if err != nil {
		// panic when error on load content
		os.Exit(1)
	}
	return string(b), nil
}
