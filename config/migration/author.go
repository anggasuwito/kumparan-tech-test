package migration

import "database/sql"

func CreateTableAuthor(db *sql.DB) error {
	var (
		query = `CREATE TABLE IF NOT EXISTS author (
						id 		   UUID PRIMARY KEY,
						created_at TIMESTAMPTZ,
						created_by VARCHAR,
						updated_at TIMESTAMPTZ,
						updated_by VARCHAR,
						deleted_at TIMESTAMPTZ,
						deleted_by VARCHAR,
						note       VARCHAR,
						name	   VARCHAR
					)`
	)

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func InitialDataAuthor(db *sql.DB) error {
	var (
		query = []string{
			`INSERT INTO author
					(id, created_at, name)
				 VALUES
				    ('825f7035-e8ea-477f-9be6-f4776ab95403', now(), 'John Wick')
				 ON CONFLICT (id) DO NOTHING`,
			`INSERT INTO author
					(id, created_at, name)
				 VALUES
				    ('e611b97d-381c-4e69-bc5f-32517c38c324', now(), 'Sarah Montana')
				 ON CONFLICT (id) DO NOTHING`,
		}
	)

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	for _, q := range query {
		_, err = tx.Exec(q)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
