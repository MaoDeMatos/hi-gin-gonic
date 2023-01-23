package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func AddDelayed(app *gin.Engine) {
	delayedRoute := app.Group("/delay")

	delayedRoute.GET("/", func(ctx *gin.Context) {
		delay := 2
		sleep(delay)
		ctx.JSON(
			http.StatusOK,
			gin.H{"data": fmt.Sprintf("You waited %d seconds", delay)},
		)
	})

	delayedRoute.GET("/:seconds", func(ctx *gin.Context) {
		delay, err := strconv.Atoi(ctx.Params.ByName("seconds"))

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		} else {
			sleep(delay)
			ctx.JSON(
				http.StatusOK,
				gin.H{"data": fmt.Sprintf("Welcome from Gin server, after %d seconds", delay)},
			)
		}
	})
}

func sleep(delay int) {
	time.Sleep(time.Duration(delay) * time.Second)
}
