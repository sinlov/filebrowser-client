## login

```go
package main
import (
	"fmt"
	"github.com/sinlov/filebrowser-client/file_browser_client"
	"testing"
)

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
```

## Get Resource

```go
package main
import (
	"fmt"
	"github.com/sinlov/filebrowser-client/file_browser_client"
	"testing"
)

func TestResourcesGet(t *testing.T) {
	if envCheck(t) {
		t.Log("must check env then test")
		return
	}
	t.Logf("~> mock ResourcesGetCheckSum")
	// mock ResourcesGetCheckSum
	client, err := tryLoginClient(t, envDebug)
	if err != nil {
		t.Fatalf("login fail!")
		return
	}
	// do ResourcesGetCheckSum
	t.Logf("~> do ResourcesGetCheckSum")
	_, err = client.ResourcesGetCheckSum("", "abc")
	if err == nil {
		t.Fatalf("client.ResourcesGetCheckSum not cover unsupport checksum")
	}
	resources, err := client.ResourcesGet("")
	// verify ResourcesGetCheckSum
	if err != nil {
		t.Fatalf("client.ResourcesGetCheckSum err: %v", err)
	}
	assert.Equal(t, `/`, resources.Path)
}
```

## post Resource and share

```go
package main
import (
	"fmt"
	"github.com/sinlov/filebrowser-client/file_browser_client"
	"testing"
)

func TestResourcesPostOne(t *testing.T) {
	// mock ResourcesPostFile
	t.Logf("~> mock ResourcesPostFile")
	t.Logf("~> do ResourcesPostFile")
	// do ResourcesPostFile
	client, err := tryLoginClient(t, envDebug)
	if err != nil {
		t.Fatalf("login fail!")
		return
	}

	localJsonFilePath := walkAllJsonFileBySuffix[len(walkAllJsonFileBySuffix)-1]
	remotePath := strings.Replace(localJsonFilePath, testDataPostFolderPath, "", -1)
	remotePath = folder.Path2WebPath(remotePath)
	var resourcePost = file_browser_client.ResourcePostFile{
		LocalFilePath:  localJsonFilePath,
		RemoteFilePath: remotePath,
	}
	err = client.ResourcesPostFile(resourcePost, true)
	// verify ResourcesPostFile
	if err != nil {
		t.Fatalf("try client.ResourcesPostFile err: %v", err)
	}

	err = client.ResourcesPostFile(resourcePost, false)
	if err == nil {
		t.Fatalf("try client.ResourcesPostFile not cover override")
	}
	t.Logf("~> do sharesRespInfinite")
	passWord := randomStr(10)
	shareResourceInfinite := file_browser_client.ShareResource{
		RemotePath: remotePath,
		ShareConfig: web_api.ShareConfig{
			Password: passWord,
			Expires:  "10",
			Unit:     web_api.ShareUnitHours,
		},
	}
	sharesRespInfinite, err := client.SharePost(shareResourceInfinite)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("------- path: %s\ndonwload page: %s \npasswd: %s", sharesRespInfinite.RemotePath, sharesRespInfinite.DownloadPage, sharesRespInfinite.DownloadPasswd)
	t.Logf("download url: %s", sharesRespInfinite.DownloadUrl)
}

```