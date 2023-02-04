package file_browser_client

import (
	"fmt"
	"github.com/monaco-io/request"
	"github.com/monaco-io/request/response"
	"github.com/sinlov/filebrowser-client/tools/folder"
	tools "github.com/sinlov/filebrowser-client/tools/str_tools"
	"github.com/sinlov/filebrowser-client/web_api"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
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

// NewClient
// new client for filebrowser
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
	return client.client(
		username, password, baseUrl,
		timeoutSecond, timeoutFileSecond,
	)
}

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

// Debug
// open FileBrowserClient debug or close
func (f *FileBrowserClient) Debug(isDebug bool) {
	f.isDebug = isDebug
}

// IsLogin
// check FileBrowserClient has login
func (f *FileBrowserClient) IsLogin() bool {
	return f.authHeadVal != ""
}

// Login
// do login by FileBrowserClient
func (f *FileBrowserClient) Login() error {
	if f.baseUrl == "" || f.username == "" {
		return fmt.Errorf("clinet not init by baseUrl or username, please check")
	}
	if f.isDebug {
		log.Printf("debug: FileBrowserClient try Login user: [ %s ] api: %s", f.username, web_api.ApiLogin())
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
	send, err := f.sendRespRaw(c, "Login", true)
	if err != nil {
		return err
	}
	if f.isDebug {
		log.Printf("debug: try Login user succes by code [ %v ]", send.Code())
	}
	f.authHeadVal = send.String()
	return nil
}

// ResourcesGet
// pathResource path resource at remote
func (f *FileBrowserClient) ResourcesGet(pathResource string) (web_api.Resources, error) {
	return f.ResourcesGetCheckSum(pathResource, "")
}

// ResourcesGetCheckSum
// pathResource path resource at remote
// checksum will be [ md5 sha1 sha256 sha512 ] or empty
func (f *FileBrowserClient) ResourcesGetCheckSum(pathResource string, checksum string) (web_api.Resources, error) {
	var resource web_api.Resources
	if !f.IsLogin() {
		if checksum == "" {
			return resource, fmt.Errorf("plase Login then get resource ResourcesGet")
		}
		return resource, fmt.Errorf("plase Login then get resource ResourcesGetCheckSum")
	}

	if !tools.StrInArr(checksum, web_api.ChecksumsDefine()) {
		return resource, fmt.Errorf("plase check checksum, now [ %s ] only support %v", checksum, web_api.ChecksumsDefine())
	}

	if pathResource != "" {
		pathResource = strings.TrimPrefix(pathResource, "/")
	}

	var urlPath = fmt.Sprintf("%s/%s", web_api.ApiResources(), pathResource)
	header := BaseHeader()
	header[web_api.AuthHeadKey] = f.authHeadVal
	var c request.Client
	if checksum == "" {
		c = request.Client{
			Timeout: time.Duration(f.timeoutSecond) * time.Second,
			URL:     urlPath,
			Method:  request.GET,
			Header:  header,
		}
	} else {
		c = request.Client{
			Timeout: time.Duration(f.timeoutSecond) * time.Second,
			URL:     urlPath,
			Method:  request.GET,
			Header:  header,
			Query: map[string]string{
				"checksum": checksum,
			},
		}
	}

	_, err := f.sendRespJson(c, &resource, "ResourcesGetCheckSum")
	if err != nil {
		return resource, err
	}

	return resource, nil
}

// ResourcesDeletePath
// remotePath just use remote path
func (f *FileBrowserClient) ResourcesDeletePath(remotePath string) (bool, error) {
	if !f.IsLogin() {
		return false, fmt.Errorf("plase Login then ResourcesDeletePath")
	}
	if remotePath == "" {
		return false, fmt.Errorf("plase check now is empty for remotePath")
	}
	remotePath = strings.TrimPrefix(remotePath, "/")
	urlPath := fmt.Sprintf("%s/%s", web_api.ApiResources(), remotePath)
	header := BaseHeader()
	header[web_api.AuthHeadKey] = f.authHeadVal
	c := request.Client{
		Timeout: time.Duration(f.timeoutFileSecond) * time.Second,
		URL:     urlPath,
		Method:  request.DELETE,
		Header:  header,
	}

	_, err := f.sendRespRaw(c, "ResourcesDeletePath", true)
	if err != nil {
		return false, fmt.Errorf("delete path error\nremote path: %s\nerr: %v", remotePath, err)
	}

	return true, nil
}

// ResourcesPostFile
// param post file by ResourcePostFile, ResourcePostFile.LocalFilePath must exist;
// override will want override remote path, but success must enable the permission at filebrowser to modify files
func (f *FileBrowserClient) ResourcesPostFile(resourceFile ResourcePostFile, override bool) error {
	if !f.IsLogin() {
		return fmt.Errorf("plase Login then ResourcesPostFile")
	}
	if resourceFile.LocalFilePath == "" {
		return fmt.Errorf("plase check LocalFilePath, now is empty for RemoteFilePath: %s", resourceFile.RemoteFilePath)
	}

	exists, err := folder.PathExists(resourceFile.LocalFilePath)
	if err != nil || !exists {
		return fmt.Errorf("plase check LocalFilePath, now is not exist at: %s , err: %v", resourceFile.LocalFilePath, err)
	}
	if folder.PathIsDir(resourceFile.LocalFilePath) {
		return fmt.Errorf("plase check LocalFilePath, now is path is folder at: %s", resourceFile.LocalFilePath)
	}

	if resourceFile.RemoteFilePath != "" {
		resourceFile.RemoteFilePath = strings.TrimPrefix(resourceFile.RemoteFilePath, "/")
	}

	urlPath := fmt.Sprintf("%s/%s", web_api.ApiResources(), resourceFile.RemoteFilePath)
	header := BaseHeader()
	header[web_api.AuthHeadKey] = f.authHeadVal
	c := request.Client{
		Timeout: time.Duration(f.timeoutFileSecond) * time.Second,
		URL:     urlPath,
		Method:  request.POST,
		Header:  header,
		Query: map[string]string{
			"override": strconv.FormatBool(override),
		},
		MultipartForm: request.MultipartForm{
			Files: []string{resourceFile.LocalFilePath},
		},
	}
	_, err = f.sendRespRaw(c, "ResourcesPost", false)
	if err != nil {
		return fmt.Errorf("post file error\nremote: %s\nlocal: %s\nerr: %v", resourceFile.RemoteFilePath, resourceFile.LocalFilePath, err)
	}

	return nil
}

type ResourcesPostDirectoryResult struct {
	FullSuccess  bool               `json:"full_success"`
	SuccessFiles []ResourcePostFile `json:"post_success_files,omitempty"`
	FailFiles    []ResourcePostFile `json:"fail_files,omitempty"`
}

// ResourcesPostDirectoryFiles
// post directory full files by ResourcePostDirectory settings
// ResourcePostDirectory.LocalDirectoryPath must exist
// override will want override remote path, but success must enable the permission at filebrowser to modify files
func (f *FileBrowserClient) ResourcesPostDirectoryFiles(resourceDirectory ResourcePostDirectory, override bool) (ResourcesPostDirectoryResult, error) {
	var result ResourcesPostDirectoryResult
	if !f.IsLogin() {
		return result, fmt.Errorf("plase Login then ResourcesPostDirectoryFiles")
	}
	if resourceDirectory.LocalDirectoryPath == "" {
		return result, fmt.Errorf("plase check LocalDirectoryPath, now is empty for RemoteDirectoryPath: %s", resourceDirectory.RemoteDirectoryPath)
	}
	exists, err := folder.PathExists(resourceDirectory.LocalDirectoryPath)
	if err != nil || !exists {
		return result, fmt.Errorf("plase check LocalDirectoryPath, now is not exist at: %s , err: %v", resourceDirectory.LocalDirectoryPath, err)
	}
	if folder.PathIsFile(resourceDirectory.LocalDirectoryPath) {
		return result, fmt.Errorf("plase check LocalDirectoryPath, now is path is file at: %s", resourceDirectory.LocalDirectoryPath)
	}
	var resourcePostFileList []ResourcePostFile
	err = filepath.Walk(resourceDirectory.LocalDirectoryPath, func(filename string, fi os.FileInfo, err error) error {
		if fi.IsDir() { // ignore dir
			return nil
		}
		innerPath := strings.Replace(filename, resourceDirectory.LocalDirectoryPath, "", -1)
		innerPath = strings.TrimPrefix(innerPath, string(filepath.Separator))
		innerPathWeb := strings.Replace(innerPath, `\`, `/`, -1)
		remoteWebPath := fmt.Sprintf("%s/%s", resourceDirectory.RemoteDirectoryPath, innerPathWeb)
		resourcePostFileList = append(resourcePostFileList, ResourcePostFile{
			LocalFilePath:  filename,
			RemoteFilePath: remoteWebPath,
		})
		return nil
	})
	if len(resourcePostFileList) == 0 {
		return result, fmt.Errorf("plase check LocalDirectoryPath, has no files at: %s", resourceDirectory.LocalDirectoryPath)
	}
	if f.isDebug {
		log.Print("debug: want ResourcesPostDirectoryFiles start\n")
		for _, resourcePostFile := range resourcePostFileList {
			log.Printf("debug: ResourcesPostDirectoryFiles\nLocalFilePath: %s\nRemoteFilePath: %s\n", resourcePostFile.LocalFilePath, resourcePostFile.RemoteFilePath)
		}
		log.Print("debug: want ResourcesPostDirectoryFiles end")
	}
	result.FullSuccess = true
	var postSuccFileList []ResourcePostFile
	var postFailFileList []ResourcePostFile
	for _, resourcePostFile := range resourcePostFileList {
		errPostFile := f.ResourcesPostFile(resourcePostFile, override)
		if errPostFile != nil {
			if f.isDebug {
				log.Printf("post folder fail at\nLocalFilePath: %s\nRemoteFilePath: %s\nerr: %s", resourcePostFile.LocalFilePath, resourcePostFile.LocalFilePath, errPostFile)
			}
			postFailFileList = append(postFailFileList, resourcePostFile)
			result.FullSuccess = false
		} else {
			postSuccFileList = append(postSuccFileList, resourcePostFile)
		}
	}
	result.FailFiles = postFailFileList
	result.SuccessFiles = postSuccFileList

	return result, nil
}

// ResourceDownload
// remotePath must exist and not empty;
// localPath must not empty and parent folder must exist
// override is overrider download
func (f *FileBrowserClient) ResourceDownload(remotePath string, localPath string, override bool) error {
	if !f.IsLogin() {
		return fmt.Errorf("plase Login then ResourceDownload")
	}
	if remotePath == "" {
		return fmt.Errorf("please check remotePath, now is empty")
	}
	remotePath = strings.TrimPrefix(remotePath, "/")

	if localPath == "" {
		return fmt.Errorf("please check localPath, now is empty")
	}

	urlPath := fmt.Sprintf("%s/%s?auth=%s", web_api.ApiRaw(), remotePath, f.authHeadVal)
	header := BaseHeader()
	header[web_api.AuthHeadKey] = f.authHeadVal
	c := request.Client{
		Timeout: time.Duration(f.timeoutFileSecond) * time.Second,
		URL:     urlPath,
		Method:  request.GET,
		Header:  header,
	}

	_, err := f.sendSaveFile(c, "ResourceDownload", localPath, override)
	if err != nil {
		return err
	}

	return nil
}

type ShareResource struct {
	RemotePath  string
	ShareConfig web_api.ShareConfig
}

type ShareContent struct {
	ShareLink      web_api.ShareLink `json:"share_link"`
	RemotePath     string            `json:"remote_path"`
	DownloadUrl    string            `json:"download_url"`
	DownloadPage   string            `json:"download_page"`
	DownloadPasswd string            `json:"page_passwd,omitempty"`
}

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
	c := request.Client{
		Timeout: time.Duration(f.timeoutSecond) * time.Second,
		URL:     urlPath,
		Method:  request.POST,
		Header:  header,
		JSON:    shareResource.ShareConfig,
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
