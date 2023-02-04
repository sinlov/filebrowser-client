package file_browser_client_test

import (
	"github.com/sinlov/filebrowser-client/file_browser_client"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckFileShaHex(t *testing.T) {
	// mock FileSha256Hex
	t.Logf("~> mock FileSha256Hex")
	// do FileSha256Hex
	t.Logf("~> do FileSha256Hex")
	fileSha256Hex, err := file_browser_client.FileSha256Hex("file_tools_test.go")
	if err != nil {
		t.Fatal(err)
	}
	// verify FileSha256Hex
	t.Logf("res FileSha256Hex fileSha256Hex: %s", fileSha256Hex)
	assert.NotEqual(t, "", fileSha256Hex)
}
