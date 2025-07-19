package config

import (
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

var (
	val = new(Configuration)
)

type Configuration struct {
	DBMaster                   *gorm.DB
	RedisClient                *redis.Client
	HttpHost                   string
	HttpPort                   string
	AppVersion                 string
	AccessTokenSecret          string
	AccessTokenExpireDuration  string
	RefreshTokenSecret         string
	RefreshTokenExpireDuration string
}

func SetConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[config] error when loading .env file " + err.Error())
	}

	dbAutoMigrate, _ := strconv.ParseBool(os.Getenv("DB_AUTO_MIGRATE"))
	dbMaster, err := getDatabase(dbConfig{
		host:        os.Getenv("DB_HOST"),
		user:        os.Getenv("DB_USER"),
		password:    os.Getenv("DB_PASSWORD"),
		dbName:      os.Getenv("DB_NAME"),
		port:        os.Getenv("DB_PORT"),
		sslMode:     os.Getenv("DB_SSL"),
		timezone:    os.Getenv("DB_TIMEZONE"),
		autoMigrate: dbAutoMigrate,
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
	val.AccessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")
	val.AccessTokenExpireDuration = os.Getenv("ACCESS_TOKEN_EXPIRE_DURATION")
	val.RefreshTokenSecret = os.Getenv("REFRESH_TOKEN_SECRET")
	val.RefreshTokenExpireDuration = os.Getenv("REFRESH_TOKEN_EXPIRE_DURATION")
}

func GetConfig() *Configuration {
	return val
}
