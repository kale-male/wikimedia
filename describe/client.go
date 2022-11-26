package describe

import (
	"fmt"
	"regexp"

	"github.com/go-resty/resty/v2"
)

type WikimediaClient interface {
	QueryText(name string) (string, error)
}

type restyWikimediaClient struct {
	client *resty.Client
}

// Helper to cover all the query params
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

type QueryResult struct {
	Query struct {
		Pages []struct {
			Revisions []struct {
				Content string
			}
		}
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

// QueryText implements WikimediaClient
func (client *restyWikimediaClient) QueryText(name string) (string, error) {
	result := new(QueryResult)
	resp, err := client.client.R().
		SetQueryParams(makeWikiQuery(name)).
		SetHeader("Accept", "application/json").
		SetResult(result).
		Get("/w/api.php")
	if err != nil {
		return "", err
	}
	result = resp.Result().(*QueryResult)

	// if len(result.Query.Pages)
	pages := result.Query.Pages
	if len(pages) == 0 || len(pages[0].Revisions) == 0 {
		return "", DescriptionNotFound{name}
	}
	description, found := parseDescription(pages[0].Revisions[0].Content)
	if !found {
		return "", DescriptionNotFound{name}
	}
	return description, nil
}

// Base client class wikipedia.org
//
// https://www.mediawiki.org/wiki/API:Main_page
func MakeHttpClient(config *Config) WikimediaClient {
	client := resty.New().EnableTrace().SetBaseURL(config.WikiAPIbaseURL)

	return &restyWikimediaClient{client}
}

type DescriptionNotFound struct {
	name string
}

// Error implements error
func (err DescriptionNotFound) Error() string {
	return fmt.Sprintf("Description for %v not found", err.name)
}
