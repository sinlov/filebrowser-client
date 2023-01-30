package web_api

import "fmt"

const (
	AuthHeadKey = "X-Auth"
	baseApi     = "api"
)

var (
	filebrowserBaseUrl = ""
	apiBaseUrl         = ""
)

func SetApiBase(baseUrl string) {
	filebrowserBaseUrl = baseUrl
	apiBaseUrl = fmt.Sprintf("%s/%s", baseUrl, baseApi)
}

func BaseUrl() string {
	if filebrowserBaseUrl == "" {
		panic("please use SetApiBase() first")
	}
	return filebrowserBaseUrl
}

func ApiBase() string {
	if apiBaseUrl == "" {
		panic("please use SetApiBase() first")
	}
	return apiBaseUrl
}
