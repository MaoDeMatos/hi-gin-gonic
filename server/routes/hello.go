package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddHello(app *gin.Engine) {
	helloRoute := app.Group("/hello")

	helloRoute.GET("/", greet)
	helloRoute.GET("/:name", greet)
}

func greet(ctx *gin.Context) {
	user := ctx.Params.ByName("name")
	if len(user) > 0 {
		ctx.JSON(http.StatusOK, gin.H{"data": user})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"data": "Hello World!"})
	}
}
