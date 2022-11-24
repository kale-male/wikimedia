package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func hello() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	}
}

type QueryResult struct {
	Query struct {
		Pages []struct {
			Revisions []struct {
				Content string
			}
		}
	}
}

func query(client *resty.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result := new(QueryResult)
		q := ctx.Query("name")
		resp, err := client.R().
			SetQueryParams(makeWikiQuery(q)).
			SetHeader("Accept", "application/json").
			SetResult(result).
			Get("https://en.wikipedia.org/w/api.php")
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		result = resp.Result().(*QueryResult)
		ctx.JSON(http.StatusOK, result)
	}
}

func makeWikiQuery(name string) map[string]string {
	return map[string]string{
		"action":        "query",
		"prop":          "revisions",
		"titles":        name,
		"rvlimit":       "1",
		"formatversion": "2",
		"format":        "json",
		"rvprop":        "content",
	}
}
