package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"test/config"
	"test/service"
)

func init() {
	config.Init()
}
func main() {
	route := gin.Default()
	route.POST("/create", service.Create1)
	route.GET("/", service.Read1)
	route.GET("/:id",service.ReadId1)
	route.DELETE("/remove/:id", service.Remove1)
	route.PUT("/update/:id", service.Update1)
	log.SetLevel(config.Config.LogLevel)
	route.Run(config.Config.Port)
}
