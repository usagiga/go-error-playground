package cockroachdb

import "github.com/usagiga/go-error-playground/commons"

type ApplicationImpl struct {
	mod SomeModule
}

func NewApplication(mod SomeModule) (app commons.Application) {
	return &ApplicationImpl{
		mod: mod,
	}
}

func (app *ApplicationImpl) Name() string {
	return "cockroachdb"
}

func (app *ApplicationImpl) Run() {
	// TODO: use module & handling errors here
	panic("implement me")
}

// TODO: Plan errors what testing
