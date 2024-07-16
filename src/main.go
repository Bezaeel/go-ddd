package main

import (
	infra "go-ddd/src/infrastructure"
	infrastructure_database "go-ddd/src/infrastructure/database"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(".env could not be loaded")
	}
	_ = godotenv.Load()

	dbHost, exist := os.LookupEnv("DB_HOST")

	if !exist {
		log.Fatal("DB_HOST not set in .env")
	}

	dbPort, exist := os.LookupEnv("DB_PORT")

	if !exist {
		log.Fatal("DB_PORT not set in .env")
	}

	dbUser, exist := os.LookupEnv("DB_USER")

	if !exist {
		log.Fatal("DB_USER not set in .env")
	}

	dbPass, exist := os.LookupEnv("DB_PASS")

	if !exist {
		log.Fatal("DB_PASS not set in .env")
	}

	dbName, exist := os.LookupEnv("DB_NAME")

	if !exist {
		log.Fatal("DB_NAME not set in .env")
	}

	infrastructure_database.ConnectToDB(dbUser, dbPass, dbHost, dbPort, dbName)
	apiServer := infra.NewAPIServer(":8000")
	apiServer.Run()
}
