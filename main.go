package main

import (
	"fmt"
	"github.com/usagiga/go-error-playground/cockroachdb"
	"github.com/usagiga/go-error-playground/commons"
	"github.com/usagiga/go-error-playground/standard"
)

func main() {
	// injection dependencies
	apps := make(map[string]commons.Application)

	sMod := standard.NewSomeModule()
	sApp := standard.NewApplication(sMod)

	cMod := cockroachdb.NewSomeModule()
	cApp := cockroachdb.NewApplication(cMod)

	apps[sApp.Name()] = sApp
	apps[cApp.Name()] = cApp

	// run apps
	config, err := readArgs()
	if err != nil {
		panic(err)
	}

	app := apps[config.AppName]
	if app == nil {
		panic(fmt.Sprintf("app '%s' not found", config.AppName))
	}

	app.Run()
}
