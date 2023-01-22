package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Create the app
var app = gin.Default()

func main() {
	// Load `.env` file
	godotenv.Load()

	/**
	* Routes
	 */

	// Ping test
	app.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"response": "pong"})
	})

	// Delayed response
	app.GET("/delay/:seconds", func(ctx *gin.Context) {
		delay, err := strconv.ParseInt(ctx.Params.ByName("seconds"), 10, 0)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		} else {
			time.Sleep(time.Duration(delay) * time.Second)
			ctx.JSON(
				http.StatusOK,
				gin.H{"data": fmt.Sprintf("Welcome Gin Server, after %d seconds", delay)},
			)
		}
	})

	helloRoute := app.Group("/hello")

	helloRoute.GET("/", greet)
	helloRoute.GET("/:name", greet)

	/**
	* Routes end
	 */

	envPort := os.Getenv("PORT")
	var port string

	if len(envPort) > 0 {
		port = ":" + envPort
	} else {
		port = ":8080"
	}

	// Listen and Server in 0.0.0.0:8080
	app.Run(port)
}

func greet(ctx *gin.Context) {
	user := ctx.Params.ByName("name")
	if len(user) > 0 {
		ctx.JSON(http.StatusOK, gin.H{"data": user})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"data": "Hello World!"})
	}
}
