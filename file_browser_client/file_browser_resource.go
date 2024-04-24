package file_browser_client

import "github.com/sinlov/filebrowser-client/web_api"

type (
	ResourcePostFile struct {
		RemoteFilePath string
		LocalFilePath  string
	}

	ResourcePostDirectory struct {
		RemoteDirectoryPath string
		LocalDirectoryPath  string
	}
)

type FileBrowserResourceFunc interface {
	ResourcesGet(pathResource string) (web_api.Resources, error)

	ResourcesGetCheckSum(pathResource string, checksum string) (web_api.Resources, error)

	ResourcesDeletePath(remotePath string) (bool, error)

	ResourcesPostFile(resourceFile ResourcePostFile, override bool) error

	ResourceDownload(remotePath string, localPath string, override bool) error
}
