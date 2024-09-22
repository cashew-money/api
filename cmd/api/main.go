package main

import (
	"flag"
	"log/slog"
	"os"
)

type Config struct {
	Port int
}

type Application struct {
	Config Config
	Logger *slog.Logger
}

func main() {
	var cfg Config

	flag.IntVar(&cfg.Port, "port", 4000, "API server port")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &Application{
		Config: cfg,
		Logger: logger,
	}

	err := app.Serve()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
