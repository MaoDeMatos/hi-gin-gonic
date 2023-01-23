package server

import (
	"hi-gin-gonic/server/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createRoutes(app *gin.Engine) {
	app.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"res": "pong"})
	})

	routes.AddDelayed(app)
	routes.AddHello(app)
}
