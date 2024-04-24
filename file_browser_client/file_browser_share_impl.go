package file_browser_client

import (
	"fmt"
	"github.com/monaco-io/request"
	"github.com/sinlov/filebrowser-client/web_api"
	"strconv"
	"strings"
	"time"
)

// SharePost
// post share by ShareResource settings
// ShareResource.RemotePath must exist
func (f *FileBrowserClient) SharePost(shareResource ShareResource) (ShareContent, error) {
	var shareContent ShareContent
	if !f.IsLogin() {
		return shareContent, fmt.Errorf("plase Login then SharePost")
	}

	if shareResource.RemotePath == "" {
		return shareContent, fmt.Errorf("please check shareResource.RemoteFilePath , now is empty")
	}

	if web_api.CheckFalseShareConfig(shareResource.ShareConfig) {
		return shareContent, fmt.Errorf("please check shareResource.ShareConfig error of setting most is Unit not in %v , or Expires is less than 0", web_api.ShareUnitDefine())
	}

	shareResource.RemotePath = strings.TrimPrefix(shareResource.RemotePath, "/")

	urlPath := fmt.Sprintf("%s/%s", web_api.ApiShare(), shareResource.RemotePath)
	header := BaseHeader()
	header[web_api.AuthHeadKey] = f.authHeadVal
	parseExpires, errParseExpires := strconv.ParseInt(shareResource.ShareConfig.Expires, 10, 0)
	if errParseExpires != nil {
		return shareContent, fmt.Errorf("please check shareResource.ShareConfig.Expires err: %v", errParseExpires)
	}
	var c request.Client
	// fix
	if parseExpires < 1 {
		c = request.Client{
			Timeout: time.Duration(f.timeoutSecond) * time.Second,
			URL:     urlPath,
			Method:  request.POST,
			Header:  header,
			JSON: web_api.ShareConfig{
				Password: shareResource.ShareConfig.Password,
			},
		}
	} else {
		c = request.Client{
			Timeout: time.Duration(f.timeoutSecond) * time.Second,
			URL:     urlPath,
			Method:  request.POST,
			Header:  header,
			JSON:    shareResource.ShareConfig,
		}
	}

	var shareLink web_api.ShareLink
	_, err := f.sendRespJson(c, &shareLink, "SharePost")
	if err != nil {
		return shareContent, err
	}
	shareContent.ShareLink = shareLink
	shareContent.RemotePath = shareResource.RemotePath
	shareContent.DownloadPage = fmt.Sprintf("%s/%s", web_api.ShareUrlBase(), shareLink.Hash)
	if shareResource.ShareConfig.Password != "" {
		shareContent.DownloadPasswd = shareResource.ShareConfig.Password
		shareContent.DownloadUrl = fmt.Sprintf("%s/%s?token=%s", web_api.ApiPublicDL(), shareLink.Hash, shareLink.Token)
	} else {
		shareContent.DownloadUrl = fmt.Sprintf("%s/%s", web_api.ApiPublicDL(), shareLink.Hash)
	}
	return shareContent, nil
}

// ShareGetByRemotePath
// get share by remote path
// return shareLink will be list
func (f *FileBrowserClient) ShareGetByRemotePath(remotePath string) ([]web_api.ShareLink, error) {
	var shareLinks []web_api.ShareLink

	if !f.IsLogin() {
		return shareLinks, fmt.Errorf("plase Login then ShareGetByRemotePath")
	}

	if remotePath == "" {
		return shareLinks, fmt.Errorf("want get remotePath is empty")
	}
	remotePath = strings.TrimPrefix(remotePath, "/")

	urlPath := fmt.Sprintf("%s/%s", web_api.ApiShare(), remotePath)
	header := BaseHeader()
	header[web_api.AuthHeadKey] = f.authHeadVal

	c := request.Client{
		Timeout: time.Duration(f.timeoutSecond) * time.Second,
		URL:     urlPath,
		Method:  request.GET,
		Header:  header,
	}
	_, err := f.sendRespJson(c, &shareLinks, "ShareGetByRemotePath")
	if err != nil {
		return shareLinks, err
	}

	return shareLinks, nil
}

// SharesGet
// get full shares by user.
// warning: do not use this api at production environment
func (f *FileBrowserClient) SharesGet() ([]web_api.ShareLink, error) {
	var shareLinks []web_api.ShareLink

	if !f.IsLogin() {
		return shareLinks, fmt.Errorf("plase Login then SharesGet")
	}
	header := BaseHeader()
	header[web_api.AuthHeadKey] = f.authHeadVal

	c := request.Client{
		Timeout: time.Duration(f.timeoutSecond) * time.Second,
		URL:     web_api.ApiShares(),
		Method:  request.GET,
		Header:  header,
	}
	_, err := f.sendRespJson(c, &shareLinks, "SharesGet")
	if err != nil {
		return shareLinks, err
	}

	return shareLinks, nil
}

// ShareDelete
// delete share by share hash
// warning: For security purposes, this api always returns correct if the permission is passed
func (f *FileBrowserClient) ShareDelete(hash string) (bool, error) {

	if !f.IsLogin() {
		return false, fmt.Errorf("plase Login then ShareDelete")
	}
	if hash == "" {
		return false, fmt.Errorf("want delete share hash is empty")
	}
	header := BaseHeader()
	header[web_api.AuthHeadKey] = f.authHeadVal
	urlPath := fmt.Sprintf("%s/%s", web_api.ApiShare(), hash)
	c := request.Client{
		Timeout: time.Duration(f.timeoutSecond) * time.Second,
		URL:     urlPath,
		Method:  request.DELETE,
		Header:  header,
	}

	_, err := f.sendRespRaw(c, "ShareDelete", true)
	if err != nil {
		return false, fmt.Errorf("delete share error\nhash: %s\nerr: %v", hash, err)
	}

	return true, nil
}
