package config

import (
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

var (
	val = new(Configuration)
)

type Configuration struct {
	DBMaster    *sql.DB
	RedisClient *redis.Client
	HttpHost    string
	HttpPort    string
	AppVersion  string
}

func SetConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[config] error when loading .env file " + err.Error())
	}

	dbMaster, err := getDatabase(dbConfig{
		host:     os.Getenv("DB_HOST"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		dbName:   os.Getenv("DB_NAME"),
		port:     os.Getenv("DB_PORT"),
		sslMode:  os.Getenv("DB_SSL"),
		timezone: os.Getenv("DB_TIMEZONE"),
	})
	if err != nil {
		log.Fatal("[config] failed connecting database " + err.Error())
	}

	redisClient := getRedis(redisConfig{
		address:  os.Getenv("REDIS_ADDR"),
		password: os.Getenv("REDIS_PASSWORD"),
	})

	val.DBMaster = dbMaster
	val.RedisClient = redisClient
	val.HttpHost = os.Getenv("HTTP_HOST")
	val.HttpPort = os.Getenv("HTTP_PORT")
	val.AppVersion = os.Getenv("APP_VERSION")
}

func GetConfig() *Configuration {
	return val
}
