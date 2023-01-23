package server

import (
	"hi-gin-gonic/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() {
	config.Init()

	// Create the app
	app := gin.Default()
	createRoutes(app)

	app.Use(handleErrors)

	// Listen and Server in 0.0.0.0:8080
	app.Run(config.Current.Port)
}

func handleErrors(ctx *gin.Context) {
	log.Printf("Total Errors -> %d", len(ctx.Errors))

	if len(ctx.Errors) <= 0 {
		ctx.Next()
		return
	}

	for _, err := range ctx.Errors {
		log.Printf("Error -> %+v\n", err)
	}

	ctx.JSON(
		http.StatusInternalServerError,
		gin.H{"error": "INTERNAL_SERVER_ERROR", "data": ctx.Errors},
	)
}
