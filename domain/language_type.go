package domain

type LanguageType int

const (
	GO LanguageType = iota
	PYTHON
	JAVA
	RUBY
	NODEJS
)

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
