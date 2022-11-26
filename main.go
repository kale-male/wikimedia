package main

import (
	"wikimedia/describe"

	_ "wikimedia/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Wikimedia Description Service
// @version         0.1
// @description     Description Service
func main() {
	app := describe.MakeApp()
	r := gin.Default()
	r.GET("hello", describe.Hello())
	r.GET("query", describe.Query(&app))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}
