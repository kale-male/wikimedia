package describe

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

// Hello godoc
// @Summary	Healthcheck
// @Description  healthcheck
// @Tags         app
// @Router       /hello [get]
func Hello() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	}
}

// Wiki Description godoc
// @Summary	Wiki Description
// @Description  getting short description by name
// @Tags         app
// @Param        name  query     string                                               true  "name to search"
// @Success      200                  string    OK
// @Failure      404                  string    NOT FOUND
// @Failure      500                  string    INTERNAL_SERVER_ERROR
// @Router       /query [get]
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
