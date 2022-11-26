package describe

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	}
}

func Query(client WikimediaClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		q := ctx.Query("name")
		result, err := client.QueryText(q)
		if err != nil {
			switch err.(type) {
			case DescriptionNotFound:
				ctx.AbortWithError(http.StatusNotFound, err)
			default:
				ctx.AbortWithError(http.StatusInternalServerError, err)
			}
		}
		ctx.JSON(http.StatusOK, result)
	}
}
