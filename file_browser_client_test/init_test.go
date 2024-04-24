package file_browser_client_test

import (
	"errors"
	"github.com/sinlov-go/unittest-kit/env_kit"
	"github.com/sinlov-go/unittest-kit/unittest_file_kit"
	"github.com/sinlov/filebrowser-client/file_browser_log"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

const (
	defTimeoutSecond     = 10
	defTimeoutFileSecond = 30

	keyEnvDebug               = "ENV_DEBUG"
	keyEnvFileBrowserBaseUrl  = "ENV_FILE_BROWSER_BASE_URL"
	keyEnvFileBrowserUserName = "ENV_FILE_BROWSER_USERNAME"
	keyEnvFileBrowserPassword = "ENV_FILE_BROWSER_PASSWORD"
)

var (
	// testBaseFolderPath
	//  test base dir will auto get by package init()
	testBaseFolderPath = ""
	testGoldenKit      *unittest_file_kit.TestGoldenKit

	envDebug    = false
	envBaseUrl  = ""
	envUserName = ""
	envPassword = ""
)

func envCheck(t *testing.T) bool {
	mustSetEnvList := []string{
		keyEnvFileBrowserBaseUrl,
		keyEnvFileBrowserUserName,
		keyEnvFileBrowserPassword,
	}
	for _, item := range mustSetEnvList {
		if os.Getenv(item) == "" {
			t.Logf("plasee set env: %s, than run test\nfull need set env %v", item, mustSetEnvList)
			return true
		}
	}

	return false
}

func init() {
	testBaseFolderPath, _ = getCurrentFolderPath()
	file_browser_log.SetLogLineDeep(2)

	testGoldenKit = unittest_file_kit.NewTestGoldenKit(testBaseFolderPath)

	envDebug = env_kit.FetchOsEnvBool(keyEnvDebug, false)
	envBaseUrl = env_kit.FetchOsEnvStr(keyEnvFileBrowserBaseUrl, "")
	envUserName = env_kit.FetchOsEnvStr(keyEnvFileBrowserUserName, "")
	envPassword = env_kit.FetchOsEnvStr(keyEnvFileBrowserPassword, "")
}

// test case file tools start

// getCurrentFolderPath can get run path this golang dir
func getCurrentFolderPath() (string, error) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("can not get current file info")
	}
	return filepath.Dir(file), nil
}

// test case file tools end
