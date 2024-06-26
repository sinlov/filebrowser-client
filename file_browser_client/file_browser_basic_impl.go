package file_browser_client

import (
	"fmt"
	"github.com/monaco-io/request"
	"github.com/sinlov/filebrowser-client/file_browser_log"
	"github.com/sinlov/filebrowser-client/web_api"
	"time"
)

// Debug
// open FileBrowserClient debug or close
func (f *FileBrowserClient) Debug(isDebug bool) {
	f.isDebug = isDebug
	if f.isDebug {
		file_browser_log.OpenDebug()
	} else {
		file_browser_log.CloseDebug()
	}
}

func (f *FileBrowserClient) GetBaseUrl() string {
	return f.baseUrl
}

func (f *FileBrowserClient) GetUsername() string {
	return f.username
}

func (f *FileBrowserClient) SetRecaptcha(recaptcha string) {
	f.recaptcha = recaptcha
}

func (f *FileBrowserClient) GetRecaptcha() string {
	return f.recaptcha
}

// IsLogin
// check FileBrowserClient has login
func (f *FileBrowserClient) IsLogin() bool {
	return f.authHeadVal != ""
}

// Login
// do login in by FileBrowserClient
func (f *FileBrowserClient) Login() error {
	if f.baseUrl == "" || f.username == "" {
		return fmt.Errorf("clinet not init by baseUrl or username, please check")
	}
	file_browser_log.Debugf(" FileBrowserClient try Login user: [ %s ] api: %s", f.username, web_api.ApiLogin())

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
	file_browser_log.Debugf(" try Login user succes by code [ %v ]", send.Code())
	f.authHeadVal = send.String()
	return nil
}
