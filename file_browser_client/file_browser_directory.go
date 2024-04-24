package file_browser_client

type ResourcesPostDirectoryResult struct {
	FullSuccess  bool               `json:"full_success"`
	SuccessFiles []ResourcePostFile `json:"post_success_files,omitempty"`
	FailFiles    []ResourcePostFile `json:"fail_files,omitempty"`
}

type FileBrowserDirectoryFunc interface {
	ResourcesPostDirectoryFiles(resourceDirectory ResourcePostDirectory, override bool) (ResourcesPostDirectoryResult, error)
}
