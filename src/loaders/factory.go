package loaders

import "github.com/khanhtc1202/chio/src"

func LoaderFactory(lang src.LanguageType) src.Loader {
	switch lang {
	case src.GO:
		return &GoFileLoader{}
	default:
		return &GoFileLoader{}
	}
}
