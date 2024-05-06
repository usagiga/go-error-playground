package main

import "flag"

type Config struct {
	AppName string
}

func readArgs() (config *Config, err error) {
	// build & parse flags
	config = new(Config)
	flag.StringVar(&config.AppName, "app", "", "app name")
	flag.Parse()

	// validation
	// do nothing, it's NOT main objective of this app. focus your business :)

	return config, nil
}
