package domain

import "io/ioutil"

type SourceFile struct {
	Path string
}

func NewSourceFile(path string) *SourceFile {
	return &SourceFile{
		Path: path,
	}
}

func (s *SourceFile) Content() string {
	b, err := ioutil.ReadFile(s.Path)
	if err != nil {
		// TODO return on err load content file
		return ""
	}
	return string(b)
}
