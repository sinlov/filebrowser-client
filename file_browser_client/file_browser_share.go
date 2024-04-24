package file_browser_client

import "github.com/sinlov/filebrowser-client/web_api"

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

type FileBrowserShareFunc interface {
	SharePost(shareResource ShareResource) (ShareContent, error)

	ShareGetByRemotePath(remotePath string) ([]web_api.ShareLink, error)

	SharesGet() ([]web_api.ShareLink, error)

	ShareDelete(hash string) (bool, error)
}
