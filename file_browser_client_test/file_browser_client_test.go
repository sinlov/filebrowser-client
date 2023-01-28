package file_browser_client_test

import (
	"fmt"
	"github.com/sinlov/filebrowser-client/file_browser_client"
	"github.com/stretchr/testify/assert"
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

	login, err := client.Login()
	if err != nil {
		t.Errorf("file_browser_client.Login() err: %v", err)
	}
	// verify Login
	assert.Equal(t, true, login)
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
	login, err := client.Login()
	if err != nil {
		return client, fmt.Errorf("file_browser_client.FileBrowserClient.Login() err: %v", err)
	}
	// verify Login
	if !login {
		return client, fmt.Errorf("file_browser_client.FileBrowserClient.Login() err: %v", "not login")
	}

	return client, nil
}

func TestResourcesGet(t *testing.T) {
	if envCheck(t) {
		t.Log("must check env then test")
		return
	}
	t.Logf("~> mock ResourcesGet")
	// mock ResourcesGet
	client, err := tryLoginClient(t, envDebug)
	if err != nil {
		t.Errorf("login fail!")
		return
	}
	// do ResourcesGet
	t.Logf("~> do ResourcesGet")
	resources, err := client.ResourcesGet("")
	// verify ResourcesGet
	if err != nil {
		t.Errorf("client.ResourcesGet err: %v", err)
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
