package test

import (
	"net/http"
	"net/url"
	"strings"
)

func GetResponse(serverURL string, path string, method string, body string, authorizationEnabled bool, authorization string) *http.Response {
	u, _ := url.ParseRequestURI(serverURL)
	u.Path = path

	r, _ := http.NewRequest(method, u.String(), strings.NewReader(body))

	if authorizationEnabled {
		r.Header.Add("authorization", "Bearer "+authorization)
	}
	r.Header.Add("content-type", "application/json")
	r.Header.Add("cache-control", "no-cache")

	client := &http.Client{}
	resp, _ := client.Do(r)

	return resp
}
