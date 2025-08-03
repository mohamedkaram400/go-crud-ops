package config

import (
	"os"
	"strconv"
)

func GetMongoURI() string {
	return os.Getenv("MONGO_URI")
}

func GetDBName() string {
	return os.Getenv("DB_NAME")
}

func GetCollectionName() string {
	return os.Getenv("COLLECTION_NAME")
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4444"
	}
	return ":" + port
}

func GetRedisHost() string {
	return os.Getenv("REDIS_HOST")
}

func GetRateNumber() int {
	rateStr := os.Getenv("RATE_NUMBER")
	rateInt, err := strconv.Atoi(rateStr)
	if err != nil {
		panic("Invalid RATE_NUMBER, must be an integer")
	}
	return rateInt
}