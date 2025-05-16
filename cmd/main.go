package main

import (
	"context"
	"fmt"
	router "myapp/internal"
	"myapp/internal/generated/db"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	// Database connection
	timeoutDuration := 5 * time.Second
	timeoutCtx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()

	pool, err := pgxpool.New(timeoutCtx, "postgres://myapp:secret@localhost:5432/myapp?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()
	q := db.New(pool)

	// Start Gin router
	router := router.SetupRouter(q)
	if err := router.Run(":8080"); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to start server: %v\n", err)
		os.Exit(1)
	}
}
