package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/plaid/plaid-go/plaid"
)

type Config struct {
	Port int
	DB   struct {
		DSN string
	}
	Plaid struct {
		ClientID string
		Secret   string
	}
}

type Application struct {
	Config      Config
	Logger      *slog.Logger
	PlaidClient *plaid.APIClient
}

func main() {
	var cfg Config

	flag.IntVar(&cfg.Port, "port", 4000, "API server port")
	flag.StringVar(&cfg.DB.DSN, "db-dsn", "", "PostgreSQL DSN")

	flag.StringVar(&cfg.Plaid.ClientID, "plaid-client-id", "", "Plaid client ID")
	flag.StringVar(&cfg.Plaid.Secret, "plaid-secret", "", "Plaid secret")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(cfg.DB.DSN)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	logger.Info("database connection pool established")

	plaidClient := createPlaidClient(cfg.Plaid.ClientID, cfg.Plaid.Secret)

	app := &Application{
		Config:      cfg,
		Logger:      logger,
		PlaidClient: plaidClient,
	}

	err = app.Serve()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func openDB(dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = pool.Ping(ctx)
	if err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}

func createPlaidClient(clientID, secret string) *plaid.APIClient {
	config := plaid.NewConfiguration()
	config.AddDefaultHeader("PLAID-CLIENT-ID", clientID)
	config.AddDefaultHeader("PLAID-SECRET", secret)
	config.UseEnvironment(plaid.Sandbox)
	return plaid.NewAPIClient(config)
}
