package main

import (
	"os"

	a "github.com/foorester/cook/internal/app"
	l "github.com/foorester/cook/internal/sys/log"
)

const (
	name     = "cook"
	logLevel = "info"
)

var (
	log = l.NewLogger(logLevel)
)

func main() {
	app, err := a.NewApp(name, log)
	if err != nil {
		log.Errorf("%s exit error: %s", name, err.Error())
		os.Exit(1)
	}

	err = app.Run()
	if err != nil {
		log.Errorf("%s exit error: %s", name, err.Error())
		os.Exit(1)
	}
}
