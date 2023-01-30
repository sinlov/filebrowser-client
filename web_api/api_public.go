package web_api

import "fmt"

func ApiPublic() string {
	return fmt.Sprintf("%s/%s", ApiBase(), "public")
}

// ApiPublicDL
// to download file for share
// if share file has password can use file_browser_client.SharesGet by password add token
func ApiPublicDL() string {
	return fmt.Sprintf("%s/%s", ApiPublic(), "dl")
}
