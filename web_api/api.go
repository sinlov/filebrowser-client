package web_api

import "fmt"

const (
	AuthHeadKey = "X-Auth"
	baseApi     = "api"
)

var (
	apiBaseUrl = ""
)

func SetApiBase(baseUrl string) {
	apiBaseUrl = fmt.Sprintf("%s/%s", baseUrl, baseApi)
}

func ApiBase() string {
	if apiBaseUrl == "" {
		panic("please use SetApiBase() first")
	}
	return apiBaseUrl
}
