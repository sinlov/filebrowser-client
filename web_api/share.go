package web_api

import "fmt"

const (
	SharePasswordHeadKey = "X-SHARE-PASSWORD"
	urlShare             = "share"
)

func ShareUrlBase() string {
	return fmt.Sprintf("%s/%s", BaseUrl(), urlShare)
}
