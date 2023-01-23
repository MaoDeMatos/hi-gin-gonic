package config

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Port    string
	GinMode string
}

var Current = Config{
	Port:    ":8080",
	GinMode: gin.DebugMode,
}

func Init() {
	envPort := os.Getenv("PORT")
	if len(envPort) > 0 {
		Current.Port = ":" + envPort
	}

	envGinMode := os.Getenv("GIN_MODE")
	if len(envGinMode) > 0 {
		Current.GinMode = strings.ToLower(envGinMode)
	}

	gin.SetMode(Current.GinMode)
}
