package main

import (
	"os"

	a "github.com/foorester/cook/internal/app"
	l "github.com/foorester/cook/internal/sys/log"
)

const (
	name     = "cook"
	env      = "ck"
	logLevel = "info"
)

var (
	log l.Logger = l.NewLogger(logLevel)
)

func main() {
	app := a.NewApp(name, env, log)

	err := app.Run()
	if err != nil {
		log.Errorf("%s exit error: %s", name, err.Error())
		os.Exit(1)
	}
}
