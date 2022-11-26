package describe

import "github.com/go-resty/resty/v2"

type WikimediaClient interface {
	QueryText(name string) (string, error)
}

type restyWikimediaClient resty.Client

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

// QueryText implements WikimediaClient
func (client *restyWikimediaClient) QueryText(name string) (string, error) {
	result := new(QueryResult)
	resp, err := (*resty.Client)(client).R().
		SetQueryParams(makeWikiQuery(name)).
		SetHeader("Accept", "application/json").
		SetResult(result).
		Get("https://en.wikipedia.org/w/api.php")
	if err != nil {
		return "", err
	}
	result = resp.Result().(*QueryResult)
	return result.Query.Pages[0].Revisions[0].Content, nil
}

// Base client class wikipedia.org
//
// https://www.mediawiki.org/wiki/API:Main_page
func MakeHttpClient() WikimediaClient {
	client := resty.New().EnableTrace()

	return (*restyWikimediaClient)(client)
}
