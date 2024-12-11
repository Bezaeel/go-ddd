package main

import (
	"fmt"
	"log/slog"
	"os"

	"api/common"
	"api/features/order"

	"infrastructure/database"

	_ "github.com/joho/godotenv/autoload"
)

var (
	dbUser = common.EnvString("DB_USER", "")
	dbPass = common.EnvString("DB_PASS", "")
	dbHost = common.EnvString("DB_HOST", "")
	dbPort = common.EnvString("DB_PORT", "")
	dbName = common.EnvString("DB_NAME", "")
	port   = common.EnvString("PORT", ":8000")
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	database.ConnectToDB(dbUser, dbPass, dbHost, dbPort, dbName)
	apiHttpServer := NewAPIServer()

	// register features
	order.RegisterModule(apiHttpServer.App(), database.DB)

	logger.Info(fmt.Sprintf("starting app on port %v", port))
	logger.Error(apiHttpServer.App().Listen(port).Error())
}
