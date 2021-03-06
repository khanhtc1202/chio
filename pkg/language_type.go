package pkg

import "strings"

type LanguageType int

const (
	GO LanguageType = iota
	PYTHON
	JAVA
	RUBY
	NODEJS
)

func ValueOfLanguage(lang string) LanguageType {
	switch strings.ToLower(lang) {
	case "go":
		return GO
	default:
		return GO
	}
}

func (l LanguageType) String() string {
	switch l {
	case GO:
		return "GO"
	case PYTHON:
		return "PYTHON"
	case JAVA:
		return "JAVA"
	case RUBY:
		return "RUBY"
	case NODEJS:
		return "NODEJS"
	default:
		return "Unknown"
	}
}

func (l LanguageType) Extension() string {
	switch l {
	case GO:
		return ".go"
	case PYTHON:
		return ".py"
	case JAVA:
		return ".java"
	case RUBY:
		return ".rb"
	case NODEJS:
		return ".js"
	default:
		return "Unknown"
	}
}
