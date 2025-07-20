package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"kumparan-tech-test/config/migration"
	"strconv"
)

type dbConfig struct {
	host        string
	user        string
	password    string
	dbName      string
	port        string
	sslMode     string
	timezone    string
	autoMigrate string
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

	if autoMigrate, _ := strconv.ParseBool(config.autoMigrate); autoMigrate {
		migrations := []error{
			migration.CreateTableArticle(db),
			migration.CreateTableAuthor(db),
			migration.InitialDataAuthor(db),
		}
		for _, e := range migrations {
			if e != nil {
				return nil, e
			}
		}
	}

	return db, nil
}
