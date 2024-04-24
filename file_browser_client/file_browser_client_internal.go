package file_browser_client

import (
	"fmt"
	"github.com/monaco-io/request"
	"github.com/monaco-io/request/response"
	"github.com/sinlov/filebrowser-client/tools/folder"
	"github.com/sinlov/filebrowser-client/web_api"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func (f *FileBrowserClient) client(
	username, password, baseUrl string,
	timeoutSecond, timeoutFileSecond uint,
) (FileBrowserClient, error) {

	var fbClient FileBrowserClient
	if baseUrl == "" {
		return fbClient, fmt.Errorf("plaese set baseUrl, now is empty")
	}
	baseUrl = strings.TrimSuffix(baseUrl, `/`)
	_, err := url.Parse(baseUrl)
	if err != nil {
		return fbClient, fmt.Errorf("client baseUrl parse err: %v", err)
	}

	fbClient = FileBrowserClient{
		username:          username,
		password:          password,
		baseUrl:           baseUrl,
		timeoutSecond:     timeoutSecond,
		timeoutFileSecond: timeoutFileSecond,
	}
	web_api.SetApiBase(baseUrl)
	return fbClient, nil
}

func (f *FileBrowserClient) sendRespRaw(c request.Client, apiName string, showCurl bool) (*response.Sugar, error) {
	if f.isDebug {
		log.Printf("debug: FileBrowserClient sendRespRaw try user: [ %s ] url: %s ", f.username, c.URL)
		if showCurl {
			c.PrintCURL()
		}
	}
	send := c.Send()
	if !send.OK() {
		return send, fmt.Errorf("try %v send user [ %v ] fail: %v", apiName, f.username, send.Error())
	}
	if send.Code() != http.StatusOK {
		return send, fmt.Errorf("try %v user [ %v ] fail: code [ %v ], msg: %v", apiName, f.username, send.Code(), send.String())
	}
	if f.isDebug {
		log.Printf("debug: sendRespRaw try %v user succes by code [ %v ], content:\n%s", apiName, send.Code(), send.String())
	}
	return send, nil
}

func (f *FileBrowserClient) sendRespJson(c request.Client, data interface{}, apiName string) (*response.Sugar, error) {
	if f.isDebug {
		log.Printf("debug: FileBrowserClient sendRespJson try user: [ %s ] url: %s ", f.username, c.URL)
		c.PrintCURL()
	}
	send := c.Send()
	if !send.OK() {
		return send, fmt.Errorf("try %v send user [ %v ] fail: %v", apiName, f.username, send.Error())
	}
	if send.Code() != http.StatusOK {
		return send, fmt.Errorf("try %v user [ %v ] fail: code [ %v ], msg: %v", apiName, f.username, send.Code(), send.String())
	}
	if f.isDebug {
		log.Printf("debug: FileBrowserClient sendRespJson try %v user succes by code [ %v ], content:\n%s", apiName, send.Code(), send.String())
	}

	resp := send.ScanJSON(data)
	if !resp.OK() {
		return resp, fmt.Errorf("try FileBrowserClient sendRespJson %v ScanJSON fail: %v", apiName, resp.Error())
	}
	return resp, nil
}

func (f *FileBrowserClient) sendSaveFile(c request.Client, apiName string, fileName string, override bool) (*response.Sugar, error) {
	if f.isDebug {
		log.Printf("debug: FileBrowserClient sendSaveFile try user: [ %s ] url: %s ", f.username, c.URL)
		c.PrintCURL()
	}
	if !override && folder.PathExistsFast(fileName) {
		return nil, fmt.Errorf("sendSaveFile not override, save path exists at: %s", fileName)
	}

	pathParent := folder.PathParent(fileName)
	if !folder.PathExistsFast(pathParent) {
		return nil, fmt.Errorf("sendSaveFile fail parent path not exists at: %s", pathParent)
	}

	send := c.Send()
	if !send.OK() {
		return send, fmt.Errorf("try %v send user [ %v ] fail: %v", apiName, f.username, send.Error())
	}
	if send.Code() != http.StatusOK {
		return send, fmt.Errorf("try %v user [ %v ] fail: code [ %v ], msg: %v", apiName, f.username, send.Code(), send.String())
	}
	if f.isDebug {
		log.Printf("debug: sendSaveFile try %v user succes by code [ %v ]", apiName, send.Code())
	}
	send.SaveToFile(fileName)
	return send, nil
}