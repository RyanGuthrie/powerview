package main

import (
	"log/slog"
	"os"
	"powerview/cmd/app/wire"
)

//go:generate wire

func main() {
	app, err := wire.CreateApp()
	if err != nil {
		slog.Error("failed creating app: %s", err)
		os.Exit(1)
	}

	if err := app.StartAndWait(); err != nil {
		slog.Error("failed starting app: %s", err)
		os.Exit(2)
	}
}
