package file_browser_client

import (
	"fmt"
	"github.com/sinlov/filebrowser-client/tools/folder"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// ResourcesPostDirectoryFiles
// post directory full files by ResourcePostDirectory settings
// ResourcePostDirectory.LocalDirectoryPath must exist
// override will want override remote path, but success must enable the permission at filebrowser to modify files
func (f *FileBrowserClient) ResourcesPostDirectoryFiles(resourceDirectory ResourcePostDirectory, override bool) (ResourcesPostDirectoryResult, error) {
	var result ResourcesPostDirectoryResult
	if !f.IsLogin() {
		return result, fmt.Errorf("plase Login then ResourcesPostDirectoryFiles")
	}
	if resourceDirectory.LocalDirectoryPath == "" {
		return result, fmt.Errorf("plase check LocalDirectoryPath, now is empty for RemoteDirectoryPath: %s", resourceDirectory.RemoteDirectoryPath)
	}
	exists, err := folder.PathExists(resourceDirectory.LocalDirectoryPath)
	if err != nil || !exists {
		return result, fmt.Errorf("plase check LocalDirectoryPath, now is not exist at: %s , err: %v", resourceDirectory.LocalDirectoryPath, err)
	}
	if folder.PathIsFile(resourceDirectory.LocalDirectoryPath) {
		return result, fmt.Errorf("plase check LocalDirectoryPath, now is path is file at: %s", resourceDirectory.LocalDirectoryPath)
	}
	var resourcePostFileList []ResourcePostFile
	_ = filepath.Walk(resourceDirectory.LocalDirectoryPath, func(filename string, fi os.FileInfo, err error) error {
		if fi.IsDir() { // ignore dir
			return nil
		}
		innerPath := strings.Replace(filename, resourceDirectory.LocalDirectoryPath, "", -1)
		innerPath = strings.TrimPrefix(innerPath, string(filepath.Separator))
		innerPathWeb := strings.Replace(innerPath, `\`, `/`, -1)
		remoteWebPath := fmt.Sprintf("%s/%s", resourceDirectory.RemoteDirectoryPath, innerPathWeb)
		resourcePostFileList = append(resourcePostFileList, ResourcePostFile{
			LocalFilePath:  filename,
			RemoteFilePath: remoteWebPath,
		})
		return nil
	})

	if len(resourcePostFileList) == 0 {
		return result, fmt.Errorf("plase check LocalDirectoryPath, has no files at: %s", resourceDirectory.LocalDirectoryPath)
	}
	if f.isDebug {
		log.Print("debug: want ResourcesPostDirectoryFiles start\n")
		for _, resourcePostFile := range resourcePostFileList {
			log.Printf("debug: ResourcesPostDirectoryFiles\nLocalFilePath: %s\nRemoteFilePath: %s\n", resourcePostFile.LocalFilePath, resourcePostFile.RemoteFilePath)
		}
		log.Print("debug: want ResourcesPostDirectoryFiles end")
	}
	result.FullSuccess = true
	var postSuccessFileList []ResourcePostFile
	var postFailFileList []ResourcePostFile
	for _, resourcePostFile := range resourcePostFileList {
		errPostFile := f.ResourcesPostFile(resourcePostFile, override)
		if errPostFile != nil {
			if f.isDebug {
				log.Printf("post folder fail at\nLocalFilePath: %s\nRemoteFilePath: %s\nerr: %s", resourcePostFile.LocalFilePath, resourcePostFile.LocalFilePath, errPostFile)
			}
			postFailFileList = append(postFailFileList, resourcePostFile)
			result.FullSuccess = false
		} else {
			postSuccessFileList = append(postSuccessFileList, resourcePostFile)
		}
	}
	result.FailFiles = postFailFileList
	result.SuccessFiles = postSuccessFileList

	return result, nil
}
