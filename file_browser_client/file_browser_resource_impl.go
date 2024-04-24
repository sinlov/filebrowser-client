package file_browser_client

import (
	"fmt"
	"github.com/monaco-io/request"
	"github.com/sinlov/filebrowser-client/tools/folder"
	tools "github.com/sinlov/filebrowser-client/tools/str_tools"
	"github.com/sinlov/filebrowser-client/web_api"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

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

	fileSha256Hex, err := FileSha256Hex(resourceFile.LocalFilePath)
	if err != nil {
		return err
	}

	if resourceFile.RemoteFilePath != "" {
		resourceFile.RemoteFilePath = strings.TrimPrefix(resourceFile.RemoteFilePath, "/")
	}

	urlPath := fmt.Sprintf("%s/%s", web_api.ApiResources(), resourceFile.RemoteFilePath)
	params := url.Values{}
	reqUrl, err := url.Parse(urlPath)
	if err != nil {
		return fmt.Errorf("url parse error\nremote: %s\nlocal: %s\nerr: %v", resourceFile.RemoteFilePath, resourceFile.LocalFilePath, err)
	}
	params.Set("override", strconv.FormatBool(override))
	reqUrl.RawQuery = params.Encode()
	urlPath = reqUrl.String()

	header := BaseHeader()
	header[web_api.AuthHeadKey] = f.authHeadVal

	fileBodyIO, err := os.Open(resourceFile.LocalFilePath)
	if err != nil {
		return fmt.Errorf("try ResourcesPostFile open error\nremote: %s\nlocal: %s\nerr: %v", resourceFile.RemoteFilePath, resourceFile.LocalFilePath, err)
	}
	defer func(fileBodyIO *os.File) {
		errFileBodyIO := fileBodyIO.Close()
		if err != nil {
			log.Fatalf("try ResourcesPostFile file IO Close err: %v", errFileBodyIO)
		}
	}(fileBodyIO)

	req, err := http.NewRequest(http.MethodPost, urlPath, fileBodyIO)
	if err != nil {
		return fmt.Errorf("try ResourcesPostFile reqesut init error\nremote: %s\nlocal: %s\nerr: %v", resourceFile.RemoteFilePath, resourceFile.LocalFilePath, err)
	}

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Content-Type", "text/plain; charset=utf-8")
	for k, v := range header {
		req.Header.Add(k, v)
	}

	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//}

	httpClient := http.Client{
		Timeout: time.Duration(f.timeoutFileSecond) * time.Second,
	}

	if f.isDebug {
		log.Printf("debug: FileBrowserClient sendFile try user: [ %s ] url: %s ", f.username, req.URL.String())
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("try ResourcesPostFile Do error\nremote: %s\nlocal: %s\nerr: %v", resourceFile.RemoteFilePath, resourceFile.LocalFilePath, err)
	}
	defer func(Body io.ReadCloser) {
		errClose := Body.Close()
		if errClose != nil {
			log.Fatalf("try ResourcesPostFile close err: %v", errClose)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("try ResourcesPostFile Do read body error\nremote: %s\nlocal: %s\nerr: %v", resourceFile.RemoteFilePath, resourceFile.LocalFilePath, err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("try ResourcesPostFile resp api error resp\ncode [ %d ]\nremote: %s\nlocal: %s\nbody:\n%s",
			resp.StatusCode,
			resourceFile.RemoteFilePath,
			resourceFile.LocalFilePath,
			string(body),
		)
	}

	fileInfo, err := f.ResourcesGetCheckSum(resourceFile.RemoteFilePath, web_api.ChecksumSha256)
	if err != nil {
		return fmt.Errorf("try ResourcesPostFile then check sum error\nremote: %s\nlocal: %s\nerr: %v", resourceFile.RemoteFilePath, resourceFile.LocalFilePath, err)
	}
	if fileInfo.Checksums == nil {
		return fmt.Errorf("try ResourcesPostFile then check sum return err not return checksums \nremote: %s\nlocal: %s", resourceFile.RemoteFilePath, resourceFile.LocalFilePath)
	}
	remoteSha256, ok := fileInfo.Checksums[web_api.ChecksumSha256]
	if !ok {
		return fmt.Errorf("try ResourcesPostFile then check sum return err not return [ %s ] in checksums \nremote: %s\nlocal: %s", web_api.ChecksumSha256, resourceFile.RemoteFilePath, resourceFile.LocalFilePath)
	}

	if fileSha256Hex != remoteSha256 {
		return fmt.Errorf("try ResourcesPostFile then check sum err not same hash [ %s ]\nremoteHash: %s\nremotePath: %s\nlocalhash: %s\nlocalPath: %s",
			web_api.ChecksumSha256, remoteSha256, resourceFile.RemoteFilePath, fileSha256Hex, resourceFile.LocalFilePath)
	}

	return nil
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
