package main

import "note/internal/api"

func main() {
	server := api.NewAPIServer(":8080")
	server.Run()
}
