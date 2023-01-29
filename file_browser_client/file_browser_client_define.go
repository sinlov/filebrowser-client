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
	"strconv"
	"strings"
	"time"
)

type (
	FileBrowserClient struct {
		isDebug bool

		Recaptcha string

		username string
		password string
		baseUrl  string

		timeoutSecond     uint
		timeoutFileSecond uint

		authHeadVal string
	}
)

func NewClient(
	username, password, baseUrl string,
	timeoutSecond, timeoutFileSecond uint,
) (FileBrowserClient, error) {
	client := FileBrowserClient{}
	if timeoutSecond < 10 {
		timeoutSecond = 10
	}

	if timeoutFileSecond < 30 {
		timeoutFileSecond = 30
	}
	return client.Client(
		username, password, baseUrl,
		timeoutSecond, timeoutFileSecond,
	)
}

func (f *FileBrowserClient) Client(
	username, password, baseUrl string,
	timeoutSecond, timeoutFileSecond uint,
) (FileBrowserClient, error) {

	var fbClient FileBrowserClient
	if baseUrl == "" {
		return fbClient, fmt.Errorf("plaese set baseUrl, now is empty")
	}
	baseUrl = strings.TrimSuffix(baseUrl, "/")
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

func (f *FileBrowserClient) Debug(isDebug bool) {
	f.isDebug = isDebug
}

func (f *FileBrowserClient) IsLogin() bool {
	return f.authHeadVal != ""
}

func (f *FileBrowserClient) Login() (bool, error) {
	if f.baseUrl == "" || f.username == "" {
		return false, fmt.Errorf("clinet not init by baseUrl or username, please check")
	}
	if f.isDebug {
		log.Printf("FileBrowserClient try Login user: [ %s ] api: %s", f.username, web_api.ApiLogin())
	}

	header := BaseHeader()

	c := request.Client{
		Timeout: time.Duration(f.timeoutSecond) * time.Second,
		URL:     web_api.ApiLogin(),
		Method:  request.POST,
		Header:  header,
		JSON: web_api.LoginRequest{
			Username: f.username,
			Password: f.password,
		},
	}
	send, err := f.sendPublic(c, "Login")
	if err != nil {
		return false, err
	}
	if f.isDebug {
		log.Printf("try Login user succes by code [ %v ]", send.Code())
	}
	f.authHeadVal = send.String()
	return true, nil
}

func (f *FileBrowserClient) sendPublic(c request.Client, apiName string) (*response.Sugar, error) {
	if f.isDebug {
		log.Printf("FileBrowserClient try ResourcesGet user: [ %s ] curl", f.username)
		c.PrintCURL()
	}
	send := c.Send()
	if !send.OK() {
		return nil, fmt.Errorf("try %v send user [ %v ] fail: %v", apiName, f.username, send.Error())
	}
	if send.Code() != http.StatusOK {
		return nil, fmt.Errorf("try %v user [ %v ] fail: code [ %v ], msg: %v", apiName, f.username, send.Code(), send.String())
	}
	if f.isDebug {
		log.Printf("try %v user succes by code [ %v ], content:\n%s", apiName, send.Code(), send.String())
	}
	return send, nil
}

func (f *FileBrowserClient) sendPublicJson(c request.Client, data interface{}, apiName string) (*response.Sugar, error) {
	send, err := f.sendPublic(c, apiName)
	if err != nil {
		return send, err
	}
	resp := send.ScanJSON(data)
	if !resp.OK() {
		return resp, fmt.Errorf("try %v ScanJSON fail: %v", apiName, resp.Error())
	}
	return resp, nil
}

// ResourcesGet
// get path resource
func (f *FileBrowserClient) ResourcesGet(pathResource string) (web_api.Resources, error) {
	var resource web_api.Resources
	if !f.IsLogin() {
		return resource, fmt.Errorf("plase Login then get resource")
	}
	urlPath := fmt.Sprintf("%s/%s", web_api.ApiResources(), pathResource)
	header := BaseHeader()
	header[web_api.AuthHeadKey] = f.authHeadVal
	c := request.Client{
		Timeout: time.Duration(f.timeoutSecond) * time.Second,
		URL:     urlPath,
		Method:  request.GET,
		Header:  header,
	}
	_, err := f.sendPublicJson(c, &resource, "ResourcesGet")
	if err != nil {
		return resource, err
	}

	return resource, nil
}

func (f *FileBrowserClient) ResourcesPostOne(resourcePost ResourcePost, override bool) (bool, error) {
	if !f.IsLogin() {
		return false, fmt.Errorf("plase Login then get resource")
	}
	if resourcePost.LocalPath == "" {
		return false, fmt.Errorf("plase check LocalPath, now is empty for RemotePath: %s", resourcePost.RemotePath)
	}

	exists, err := folder.PathExists(resourcePost.LocalPath)
	if err != nil || !exists {
		return false, fmt.Errorf("plase check LocalPath, now is not exist at: %s , err: %v", resourcePost.LocalPath, err)
	}
	if folder.PathIsDir(resourcePost.LocalPath) {
		return false, fmt.Errorf("plase check LocalPath, now is path is folder at: %s", resourcePost.LocalPath)
	}

	urlPath := fmt.Sprintf("%s/%s", web_api.ApiResources(), resourcePost.RemotePath)
	header := BaseHeader()
	header[web_api.AuthHeadKey] = f.authHeadVal
	c := request.Client{
		Timeout: time.Duration(f.timeoutSecond) * time.Second,
		URL:     urlPath,
		Method:  request.POST,
		Header:  header,
		Query: map[string]string{
			"override": strconv.FormatBool(override),
		},
		MultipartForm: request.MultipartForm{
			Files: []string{resourcePost.LocalPath},
		},
	}
	_, err = f.sendPublic(c, "ResourcesPost")
	if err != nil {
		return false, fmt.Errorf("post file error\nremote: %s\nlocal: %s\nerr: %v", resourcePost.RemotePath, resourcePost.LocalPath, err)
	}

	return true, nil
}
