package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type dbConfig struct {
	host     string
	user     string
	password string
	dbName   string
	port     string
	sslMode  string
	timezone string
}

func getDatabase(config dbConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		config.host,
		config.user,
		config.password,
		config.dbName,
		config.port,
		config.timezone,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
