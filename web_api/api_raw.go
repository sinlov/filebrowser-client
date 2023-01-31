package web_api

import "fmt"

func ApiRaw() string {
	return fmt.Sprintf("%s/%s", ApiBase(), "raw")
}
