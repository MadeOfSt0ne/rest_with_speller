package main

import (
	"note/internal/api"
	"note/internal/db"
	"os"

	"github.com/lpernett/godotenv"
	"github.com/sirupsen/logrus"
	_ "modernc.org/sqlite"
)

func main() {
	loadEnv()
	db := db.ConnectDB()
	server := api.NewAPIServer(os.Getenv("SERVER_PORT"), db)
	server.Run()
	defer db.Close()
}

func loadEnv() {
	logrus.Info("loading env")
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Info("failed to load env file: ", err)
	}
}
