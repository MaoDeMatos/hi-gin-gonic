package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Create the app
var app = gin.Default()

func main() {
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

	// Listen and Server in 0.0.0.0:8080
	app.Run(":8080")
}

func greet(ctx *gin.Context) {
	user := ctx.Params.ByName("name")
	if len(user) > 0 {
		ctx.JSON(http.StatusOK, gin.H{"data": user})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"data": "Hello World!"})
	}
}
