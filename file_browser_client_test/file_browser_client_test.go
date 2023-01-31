package file_browser_client_test

import (
	"fmt"
	"github.com/sinlov/filebrowser-client/file_browser_client"
	"github.com/sinlov/filebrowser-client/tools/folder"
	"github.com/sinlov/filebrowser-client/web_api"
	"github.com/stretchr/testify/assert"
	"path"
	"strings"
	"testing"
)

func Test_NewClient(t *testing.T) {
	// mock _NewClient

	t.Logf("~> mock _NewClient")

	// do _NewClient
	t.Logf("~> do _NewClient")
	client, err := file_browser_client.NewClient(
		envUserName,
		envPassword,
		envBaseUrl,
		defTimeoutSecond,
		defTimeoutFileSecond,
	)
	if err != nil {
		if envCheck(t) {
			return
		}
		t.Errorf("file_browser_client.NewClient err: %v", err)
	}
	// verify _NewClient
	assert.Equal(t, "", client.Recaptcha)
	assert.False(t, client.IsLogin())
}

func TestLogin(t *testing.T) {
	// mock Login

	t.Logf("~> mock Login")
	client, err := file_browser_client.NewClient(
		envUserName,
		envPassword,
		envBaseUrl,
		defTimeoutSecond,
		defTimeoutFileSecond,
	)
	if err != nil {
		if envCheck(t) {
			return
		}
		t.Errorf("file_browser_client.NewClient() err: %v", err)
	}
	// do Login
	t.Logf("~> do Login")
	client.Debug(envDebug)

	assert.False(t, client.IsLogin())

	err = client.Login()
	if err != nil {
		t.Errorf("file_browser_client.Login() err: %v", err)
	}
	// verify Login
	assert.True(t, client.IsLogin())
}

func tryLoginClient(t *testing.T, isDebug bool) (file_browser_client.FileBrowserClient, error) {
	client, err := file_browser_client.NewClient(
		envUserName,
		envPassword,
		envBaseUrl,
		defTimeoutSecond,
		defTimeoutFileSecond,
	)

	if err != nil {
		return client, err
	}

	client.Debug(isDebug)
	err = client.Login()
	if err != nil {
		return client, fmt.Errorf("file_browser_client.FileBrowserClient.Login() err: %v", err)
	}
	// verify Login
	if !client.IsLogin() {
		return client, fmt.Errorf("file_browser_client.FileBrowserClient.Login() err: %v", "not login")
	}

	return client, nil
}

func TestResourcesGet(t *testing.T) {
	if envCheck(t) {
		t.Log("must check env then test")
		return
	}
	t.Logf("~> mock ResourcesGetCheckSum")
	// mock ResourcesGetCheckSum
	client, err := tryLoginClient(t, envDebug)
	if err != nil {
		t.Errorf("login fail!")
		return
	}
	// do ResourcesGetCheckSum
	t.Logf("~> do ResourcesGetCheckSum")
	_, err = client.ResourcesGetCheckSum("", "abc")
	if err == nil {
		t.Errorf("client.ResourcesGetCheckSum not cover unsupport checksum")
	}
	resources, err := client.ResourcesGet("")
	// verify ResourcesGetCheckSum
	if err != nil {
		t.Errorf("client.ResourcesGetCheckSum err: %v", err)
	}
	assert.Equal(t, "/", resources.Path)
}

func TestResourceGet_Not_Found(t *testing.T) {
	if envCheck(t) {
		t.Log("must check env then test")
		return
	}
	// mock ResourceGet_Not_Found
	client, err := tryLoginClient(t, envDebug)
	if err != nil {
		t.Errorf("login fail!")
		return
	}
	t.Logf("~> mock ResourceGet_Not_Found")
	// do ResourceGet_Not_Found
	resources, err := client.ResourcesGet("/abc")
	t.Logf("~> do ResourceGet_Not_Found")
	if err == nil {
		t.Errorf("must test not found err: %v", err)
	}
	t.Logf("mock ResourceGet_Not_Found err: %v", err)
	// verify ResourceGet_Not_Found
	assert.Empty(t, resources)
}

func TestResourcesPostOne(t *testing.T) {
	if envCheck(t) {
		t.Log("must check env then test")
		return
	}
	// mock ResourcesPostFile
	t.Logf("~> mock ResourcesPostFile")
	testDataPostFolderPath, err := initTestDataPostFileDir()
	if err != nil {
		t.Error(err)
	}

	walkAllJsonFileBySuffix, err := folder.WalkAllFileBySuffix(testDataPostFolderPath, "json")
	if err != nil {
		t.Error(err)
	}
	if len(walkAllJsonFileBySuffix) == 0 {
		t.Fatalf("walkAllJsonFileBySuffix len is 0")
	}

	t.Logf("~> do ResourcesPostFile")
	// do ResourcesPostFile
	client, err := tryLoginClient(t, envDebug)
	if err != nil {
		t.Errorf("login fail!")
		return
	}

	localJsonFilePath := walkAllJsonFileBySuffix[len(walkAllJsonFileBySuffix)-1]
	remotePath := strings.Replace(localJsonFilePath, testDataPostFolderPath, "", -1)
	remotePath = strings.TrimPrefix(remotePath, "/")
	var resourcePost = file_browser_client.ResourcePostFile{
		LocalFilePath:  localJsonFilePath,
		RemoteFilePath: remotePath,
	}
	err = client.ResourcesPostFile(resourcePost, true)
	// verify ResourcesPostFile
	if err != nil {
		t.Errorf("try client.ResourcesPostFile err: %v", err)
	}

	err = client.ResourcesPostFile(resourcePost, false)
	if err == nil {
		t.Errorf("try client.ResourcesPostFile not cover override")
	}

}

func TestSharesPost(t *testing.T) {
	// mock SharePost
	t.Logf("~> mock SharePost")
	if envCheck(t) {
		t.Log("must check env then test")
		return
	}
	// mock ResourcesPostFile
	t.Logf("~> mock ResourcesPostFile")
	testPostDataFolderPath, err := initTestDataPostFileDir()
	if err != nil {
		t.Error(err)
	}
	testDataDownloadFolderPath, err := initTestDataDownloadDir()
	if err != nil {
		t.Error(err)
	}

	walkAllJsonFileBySuffix, err := folder.WalkAllFileBySuffix(testPostDataFolderPath, "json")
	if err != nil {
		t.Error(err)
	}
	if len(walkAllJsonFileBySuffix) == 0 {
		t.Fatalf("walkAllJsonFileBySuffix len is 0")
	}

	localJsonFilePath := walkAllJsonFileBySuffix[len(walkAllJsonFileBySuffix)-1]
	remotePath := strings.Replace(localJsonFilePath, testPostDataFolderPath, "", -1)
	remotePath = strings.TrimPrefix(remotePath, "/")

	// do SharePost

	client, err := tryLoginClient(t, envDebug)
	if err != nil {
		t.Errorf("login fail!")
		return
	}

	var resourcePost = file_browser_client.ResourcePostFile{
		LocalFilePath:  localJsonFilePath,
		RemoteFilePath: remotePath,
	}
	err = client.ResourcesPostFile(resourcePost, true)
	if err != nil {
		t.Errorf("try client.ResourcesPostFile err: %v", err)
	}

	remotePathGetCheckSum256, err := client.ResourcesGetCheckSum(remotePath, web_api.ChecksumSha256)
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, remotePathGetCheckSum256.Checksums)

	downloadLocalPath := path.Join(testDataDownloadFolderPath, remotePath)
	t.Logf("downloadLocalPath: %s", downloadLocalPath)
	pathParent := folder.PathParent(downloadLocalPath)
	_ = folder.RmDirForce(pathParent)

	err = client.ResourceDownload(remotePath, downloadLocalPath, true)
	if err == nil {
		t.Error("not cover ResourceDownload not init parent path")
	}
	err = folder.Mkdir(pathParent)
	if err != nil {
		t.Error(err)
	}
	err = client.ResourceDownload(remotePath, downloadLocalPath, true)
	if err != nil {
		t.Error(err)
	}

	err = client.ResourceDownload(remotePath, downloadLocalPath, false)
	if err == nil {
		t.Error("not cover ResourceDownload not override path")
	}

	err = client.ResourceDownload(remotePath, downloadLocalPath, true)
	if err != nil {
		t.Error(err)
	}

	t.Logf("~> do SharePost")
	passWord := randomStr(10)
	shareResource := file_browser_client.ShareResource{
		RemotePath: remotePath,
		ShareConfig: web_api.ShareConfig{
			Password: passWord,
			Expires:  "10",
			Unit:     web_api.ShareUnitHours,
		},
	}
	sharesResp, err := client.SharePost(shareResource)
	if err != nil {
		t.Error(err)
	}
	// verify SharePost
	assert.NotNil(t, sharesResp)
	assert.NotEqual(t, "", sharesResp.ShareLink.Hash)
	t.Logf("------- path: %s\ndonwload page: %s \npasswd: %s", sharesResp.RemotePath, sharesResp.DownloadPage, sharesResp.DownloadPasswd)
	t.Logf("download url: %s", sharesResp.DownloadUrl)

	_, err = client.ShareGetByRemotePath("")
	if err == nil {
		t.Error("not cover ShareGetByRemotePath want delete hash is empty")
	}
	shareGetByRemotePath, err := client.ShareGetByRemotePath(remotePath)
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, shareGetByRemotePath)

	shareLinks, err := client.SharesGet()
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, shareLinks)

	_, err = client.ShareDelete("")
	if err == nil {
		t.Error("not cover ShareDelete want delete hash is empty")
	}
	shareHashMockFail := sharesResp.ShareLink.Hash + "xxx"
	shareDeleteResp, err := client.ShareDelete(shareHashMockFail)
	if err != nil {
		t.Error("not cover each hash delete can be guessed")
	}
	assert.True(t, shareDeleteResp)

	//_, err = client.ShareDelete(sharesResp.Hash)
	//if err != nil {
	//	t.Error(err)
	//}

	deletePathRes, err := client.ResourcesDeletePath(remotePath)

	if err != nil {
		t.Error(err)
	}

	assert.True(t, deletePathRes)

	err = client.ResourceDownload(remotePath, downloadLocalPath, true)
	if err == nil {
		t.Error("not cover client.ResourceDownload 404 not found")
	}

	var remoteDirPath = "inner_1"
	var resourceDirectory = file_browser_client.ResourcePostDirectory{
		LocalDirectoryPath:  path.Join(testPostDataFolderPath, remoteDirPath),
		RemoteDirectoryPath: remoteDirPath,
	}
	postDirectoryFilesRes, err := client.ResourcesPostDirectoryFiles(resourceDirectory, true)
	if err != nil {
		t.Error(err)
	}
	assert.True(t, postDirectoryFilesRes.FullSuccess)
	assert.Greater(t, len(postDirectoryFilesRes.SuccessFiles), len(postDirectoryFilesRes.FailFiles))
}
