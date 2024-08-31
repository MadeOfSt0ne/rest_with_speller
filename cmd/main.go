package main

import (
	"note/internal/api"
	"note/internal/db"

	_ "modernc.org/sqlite"
)

func main() {
	db := db.ConnectDB()
	server := api.NewAPIServer(":8080", db)
	server.Run()
	defer db.Close()
}
