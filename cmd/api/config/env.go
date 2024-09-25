package config

import (
	"context"
	"log/slog"
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

type Env struct {
	Config      Config
	Logger      *slog.Logger
	DB          *pgxpool.Pool
	PlaidClient *plaid.APIClient
}

func NewDB(dsn string) (*pgxpool.Pool, error) {
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

func NewPlaidClient(clientID, secret string) *plaid.APIClient {
	config := plaid.NewConfiguration()

	config.AddDefaultHeader("PLAID-CLIENT-ID", clientID)
	config.AddDefaultHeader("PLAID-SECRET", secret)
	config.UseEnvironment(plaid.Sandbox)

	return plaid.NewAPIClient(config)
}
