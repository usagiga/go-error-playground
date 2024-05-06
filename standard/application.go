package standard

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
	return "standard"
}

func (app *ApplicationImpl) Run() {
	// TODO: use module & handling errors here
	panic("implement me")
}

// TODO: Rewrite plans

// Go 1.13 errors
// As, Is, fmt.Errorf

// Go 1.20 errors
// Join, fmt.Errorf, Unwrap, As, Is
