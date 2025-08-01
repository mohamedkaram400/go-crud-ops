package config

import (
	"os"
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
