package server

import (
	"hi-gin-gonic/config"
)

func Init() {
	config.Init()
	router := newRouter()
	// Listen and Server in 0.0.0.0:8080
	router.Run(config.Current.Port)
}
