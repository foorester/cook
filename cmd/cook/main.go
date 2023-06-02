package main

import (
	"os"

	a "github.com/foorester/cook/internal/app"
	l "github.com/foorester/cook/internal/infra/log"
)

const (
	name = "cook"
	env  = "ck"
)

var (
	log l.Logger = l.NewLogger(l.Level.Info, false)
)

func main() {
	app := a.NewApp(name, env, log)

	err := app.Run()
	if err != nil {
		log.Errorf("%s exit error: %s", name, err.Error())
		os.Exit(1)
	}
}
