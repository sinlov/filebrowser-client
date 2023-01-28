package file_browser_client_test

import (
	"os"
	"testing"
)

const (
	defTimeoutSecond     = 10
	defTimeoutFileSecond = 30
)

var (
	envDebug    = false
	envUserName = ""
	envPassword = ""
	envBaseUrl  = ""
)

func init() {
	envDebug = os.Getenv("ENV_DEBUG") == "true"

	envUserName = os.Getenv("ENV_FILE_BROWSER_USERNAME")
	envPassword = os.Getenv("ENV_FILE_BROWSER_PASSWORD")
	envBaseUrl = os.Getenv("ENV_FILE_BROWSER_BASE_URL")
}

func envCheck(t *testing.T) bool {
	mustSetEnvList := []string{
		"ENV_FILE_BROWSER_BASE_URL",
		"ENV_FILE_BROWSER_USERNAME",
		"ENV_FILE_BROWSER_PASSWORD",
	}
	for _, item := range mustSetEnvList {
		if os.Getenv(item) == "" {
			t.Logf("plasee set env: %v, than run test", mustSetEnvList)
			return true
		}
	}
	if os.Getenv("ENV_FILE_BROWSER_BASE_URL") == "" {

		return true
	}

	return false
}
