package bootstrap

import (
	"github.com/hxzhouh/go-zen.git/storage"
	"log/slog"
	"os"
)

type Application struct {
	Env *Env
}

func App(configFile string) *Application {
	app := &Application{}
	app.Env = NewEnv(configFile)
	dsn := app.getStorageDsn()
	if dsn == "" {
		slog.Info("No storage dsn found")
		os.Exit(1)
	}
	var err error
	err = storage.InitStorage(app.Env.DBType, dsn)
	if err != nil {
		slog.Error("Failed to connect database")
		os.Exit(1)
	}
	return app
}

func (app *Application) getStorageDsn() string {
	switch app.Env.DBType {
	case "sqlite3":
		slog.Info("use sqlite3")
		return app.Env.DBHost
	}
	return ""
}

func (app *Application) Close() {
	slog.Info("Close the application")
}
