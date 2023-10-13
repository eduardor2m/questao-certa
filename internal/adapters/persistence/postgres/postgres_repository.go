package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type connectorManager interface {
	getConnection() (*sqlx.DB, error)
	closeConnection(conn *sqlx.DB)
}

var _ connectorManager = (*DatabaseConnectorManager)(nil)

type DatabaseConnectorManager struct{}

func RunMigrations() error {
	var (
		dbUser     string
		dbPassword string
		dbHost     string
		dbPort     string
		dbName     string
	)

	development := os.Getenv("DEVELOPMENT")

	if development == "true" {
		dbUser = os.Getenv("POSTGRES_USER")
		dbPassword = os.Getenv("POSTGRES_PASSWORD")
		dbHost = os.Getenv("POSTGRES_HOST")
		dbPort = os.Getenv("POSTGRES_PORT")
		dbName = os.Getenv("POSTGRES_DB")
	}

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS "user" (
		id UUID PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		admin BOOLEAN NOT NULL DEFAULT FALSE,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW()
	);
`)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	return nil
}

func (dcm *DatabaseConnectorManager) getConnection() (*sqlx.DB, error) {
	var (
		dbUser     = os.Getenv("POSTGRES_USER")
		dbPassword = os.Getenv("POSTGRES_PASSWORD")
		dbHost     = os.Getenv("POSTGRES_HOST")
		dbPort     = os.Getenv("POSTGRES_PORT")
		dbName     = os.Getenv("POSTGRES_DB")
	)

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sqlx.Connect("postgres", dbURL)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)

	return db, nil
}

func (dcm *DatabaseConnectorManager) closeConnection(conn *sqlx.DB) {
	err := conn.Close()

	if err != nil {
		panic(err)
	}
}
