package main

import (
	"hi-gin-gonic/server"

	"github.com/joho/godotenv"
)

func main() {
	// Load `.env` file
	godotenv.Load()
	// Start server
	server.Init()
}
