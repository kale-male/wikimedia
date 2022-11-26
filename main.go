package main

import (
	"wikimedia/describe"

	"github.com/gin-gonic/gin"
)

func main() {
	app := describe.MakeApp()
	r := gin.Default()
	r.GET("hello", describe.Hello())
	r.GET("query", describe.Query(&app))
	r.Run()
}
