package file_browser_client

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
