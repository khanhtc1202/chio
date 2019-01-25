package loaders

import "github.com/khanhtc1202/chio/pkg"

func LoaderFactory(lang pkg.LanguageType) pkg.Loader {
	switch lang {
	case pkg.GO:
		return &GoFileLoader{}
	default:
		return &GoFileLoader{}
	}
}
