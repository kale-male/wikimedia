package describe

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func Hello() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	}
}

func Query(app *App) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		q := ctx.Query("name")
		dlCtx, _ := context.WithTimeout(ctx, app.Config.RequestTimeout)
		result, err := app.WikimediaClient.QueryText(dlCtx, q)
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
