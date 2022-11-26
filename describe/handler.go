package describe

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func Hello() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	}
}

var descriptionRegexp = regexp.MustCompile(`(?i)\{\{Short description\|(.*)\}\}`)

func parseDescription(queryResult string) (string, bool) {
	result := descriptionRegexp.FindStringSubmatch(queryResult)
	if result == nil || len(result) < 2 {
		return "", false
	}
	return result[1], true
}

func Query(client WikimediaClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		q := ctx.Query("name")
		resp, err := client.QueryText(q)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		result, found := parseDescription(resp)
		if !found {
			ctx.Status(http.StatusNotFound)
			return
		}
		ctx.JSON(http.StatusOK, result)
	}
}
