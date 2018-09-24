package domain

type ModuleProperties interface {
	GetFanInDepend() int
	GetFanOutDepend() int
	GetAbstractMember() int
	GetConcreteMember() int
}

func FactoryModuleProperties(lang LanguageType) ModuleProperties {
	// TODO create ModuleProperties for each type of language
	switch lang {
	case GO:
		return nil
	case JAVA:
		return nil
	default:
		return nil
	}
}
