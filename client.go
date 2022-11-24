package main

import "github.com/go-resty/resty/v2"

func MakeHttpClient() *resty.Client {
	client := resty.New().EnableTrace()

	return client
}
