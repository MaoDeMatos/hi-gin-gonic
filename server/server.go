package server

import (
	"hi-gin-gonic/config"

	"github.com/gin-gonic/gin"
)

func Init() {
	config.Init()

	// Create the app
	app := gin.Default()
	createRoutes(app)

	// Listen and Server in 0.0.0.0:8080
	app.Run(config.Current.Port)
}
