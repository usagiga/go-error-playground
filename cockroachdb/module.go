package cockroachdb

type SomeModule interface {
}

type SomeModuleImpl struct {
}

func NewSomeModule() SomeModule {
	return &SomeModuleImpl{}
}
