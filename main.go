package main

import "github.com/gin-gonic/gin"

func main() {
	client := MakeHttpClient()
	r := gin.Default()
	r.GET("hello", hello())
	r.GET("query", query(client))
	r.Run()
}
