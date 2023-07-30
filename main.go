package main

import (
	"embed"
	"os"

	a "github.com/foorester/cook/internal/app"
	l "github.com/foorester/cook/internal/sys/log"
)

const (
	name     = "cook"
	logLevel = "debug"
)

var (
	//go:embed all:assets/migrations/pg/*.sql
	migFs embed.FS

	//go:embed all:assets/seeding/pg/*.sql
	seedFs embed.FS
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

	app.SetMigrationsFs(migFs)
	app.SetSeedingFs(seedFs)

	err = app.Run()
	if err != nil {
		log.Errorf("%s exit error: %s", name, err.Error())
		os.Exit(1)
	}
}
