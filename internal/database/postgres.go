package database

import (
	"database/sql"
	"log/slog"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

func NewDBConnection() (*sql.DB, error) {
	dbURI := os.Getenv("DATABASE_URL")
	db, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	slog.Info("connected to database")
	return db, nil
}
