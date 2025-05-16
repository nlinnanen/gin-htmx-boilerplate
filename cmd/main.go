package main

import (
	"context"
	"fmt"
	router "myapp/internal"
	"myapp/internal/config"
	"myapp/internal/generated/db"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to load configuration: %v\n", err)
		os.Exit(1)
	}

	timeoutDuration := 5 * time.Second
	timeoutCtx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()

	pool, err := pgxpool.New(timeoutCtx, cfg.DatabaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()
	q := db.New(pool)

	router := router.SetupRouter(q, cfg)
	if err := router.Run(":" + cfg.Port); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to start server: %v\n", err)
		os.Exit(1)
	}
}
