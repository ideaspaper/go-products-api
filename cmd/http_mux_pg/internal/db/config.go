package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	DB_HOST = "localhost"
	DB_PORT = "15432"
	DB_USER = "postgres"
	DB_PASS = "postgres"
	DB_NAME = "products_api_db"
	DB_SSLM = "disable"
)

var pool *pgxpool.Pool

func Connect() error {
	dsn := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		DB_HOST,
		DB_PORT,
		DB_USER,
		DB_PASS,
		DB_NAME,
		DB_SSLM,
	)
	pgxPool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return err
	}
	pool = pgxPool
	return nil
}

func Get() *pgxpool.Pool {
	return pool
}
