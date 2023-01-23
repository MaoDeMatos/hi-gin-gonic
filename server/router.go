package server

import (
	"fmt"
	"hi-gin-gonic/server/routes"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func newRouter() *gin.Engine {
	// Create the app
	app := gin.Default()

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

	routes.AddHello(app)

	return app
}
