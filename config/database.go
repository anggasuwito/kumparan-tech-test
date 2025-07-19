package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"kumparan-tech-test/internal/domain/model"
)

type dbConfig struct {
	host        string
	user        string
	password    string
	dbName      string
	port        string
	sslMode     string
	timezone    string
	autoMigrate bool
}

func getDatabase(config dbConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		config.host,
		config.user,
		config.password,
		config.dbName,
		config.port,
		config.timezone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if config.autoMigrate {
		err = autoMigrate(db)
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}

func autoMigrate(db *gorm.DB) error {
	err := []error{
		db.AutoMigrate(&model.Concert{}),
		db.AutoMigrate(&model.ConcertPurchaseHistory{}),
	}
	for _, e := range err {
		if e != nil {
			return e
		}
	}
	return nil
}
