package migration

import (
	"database/sql"
)

func CreateTableArticle(db *sql.DB) error {
	var (
		query = `CREATE TABLE IF NOT EXISTS article (
						id 		   UUID PRIMARY KEY,
						created_at TIMESTAMPTZ,
						created_by VARCHAR,
						updated_at TIMESTAMPTZ,
						updated_by VARCHAR,
						deleted_at TIMESTAMPTZ,
						deleted_by VARCHAR,
						note       VARCHAR,
						author_id  VARCHAR,
						title      VARCHAR,
						body       TEXT
					)`
	)
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
