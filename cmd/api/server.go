package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func (app *Application) Serve() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Config.Port),
		Handler:      routes(app),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(app.Logger.Handler(), slog.LevelError),
	}

	app.Logger.Info("starting server", "addr", srv.Addr)

	return srv.ListenAndServe()
}
