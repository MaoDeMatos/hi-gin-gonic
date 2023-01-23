package config

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Port string
}

var Current = Config{
	Port: ":8080",
}

func Init() {
	envPort := os.Getenv("PORT")
	if len(envPort) > 0 {
		Current.Port = ":" + envPort
	}

	envGinMode := strings.ToLower(os.Getenv("GIN_MODE"))
	if len(envGinMode) > 0 {
		gin.SetMode(envGinMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}
