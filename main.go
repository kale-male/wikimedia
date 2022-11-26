package main

import (
	"wikimedia/describe"

	"github.com/gin-gonic/gin"
)

func main() {
	client := describe.MakeCachedWikimediaClient(describe.MakeHttpClient())
	r := gin.Default()
	r.GET("hello", describe.Hello())
	r.GET("query", describe.Query(client))
	r.Run()
}
