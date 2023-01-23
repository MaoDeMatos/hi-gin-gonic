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

	app.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusNotFound,
			gin.H{"error": "NOT_FOUND", "message": "This endpoint does not exist"},
		)
	})

	app.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusMethodNotAllowed,
			gin.H{"error": "METHOD_NOT_ALLOWED", "message": "This method is not allowed on this endpoint"},
		)
	})
}
