package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/cashew-money/api/cmd/api/config"
)

func main() {
	var cfg config.Config

	flag.IntVar(&cfg.Port, "port", 4000, "API server port")
	flag.StringVar(&cfg.DB.DSN, "db-dsn", "", "PostgreSQL DSN")

	flag.StringVar(&cfg.Plaid.ClientID, "plaid-client-id", "", "Plaid client ID")
	flag.StringVar(&cfg.Plaid.Secret, "plaid-secret", "", "Plaid secret")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := config.NewDB(cfg.DB.DSN)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	logger.Info("database connection pool established")

	plaidClient := config.NewPlaidClient(cfg.Plaid.ClientID, cfg.Plaid.Secret)

	env := &config.Env{
		Config:      cfg,
		Logger:      logger,
		DB:          db,
		PlaidClient: plaidClient,
	}

	err = serve(env)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
