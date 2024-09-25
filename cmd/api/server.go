package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/cashew-money/api/cmd/api/config"
)

func serve(env *config.Env) error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", env.Config.Port),
		Handler:      routes(env),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(env.Logger.Handler(), slog.LevelError),
	}

	env.Logger.Info("starting server", "addr", srv.Addr)

	return srv.ListenAndServe()
}
