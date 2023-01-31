package web_api

import (
	"fmt"
	"os"
)

const (
	ChecksumMd5    = "md5"
	ChecksumSha1   = "sha1"
	ChecksumSha256 = "sha256"
	ChecksumSha512 = "sha512"
)

var defineChecksums []string

func ChecksumsDefine() []string {
	if defineChecksums == nil {
		defineChecksums = []string{
			"",
			ChecksumMd5,
			ChecksumSha1,
			ChecksumSha256,
			ChecksumSha512,
		}
	}
	return defineChecksums
}

func ApiResources() string {
	return fmt.Sprintf("%s/%s", ApiBase(), "resources")
}

type FileInfo struct {
	// at file browser path
	Path string `json:"path"`
	// Name file name
	Name string `json:"name"`
	// unit byte
	Size int64 `json:"size"`
	// Extension name of extension
	Extension string `json:"extension"`
	// Modified time utc string
	Modified string `json:"modified"`
	//
	Mode os.FileMode `json:"mode"`
	// Type [ blob text ] text will send content
	Type      string   `json:"type"`
	Subtitles []string `json:"subtitles,omitempty"`
	Content   string   `json:"content,omitempty"`
	// Checksums
	// key will be [ md5 sha1 sha256 sha512 ]
	Checksums map[string]string `json:"checksums,omitempty"`
	Token     string            `json:"token,omitempty"`
}

type ResourcesSorting struct {
	By  string `json:"by"`
	Asc bool   `json:"asc"`
}

type Resources struct {
	Sorting ResourcesSorting `json:"sorting"`
	// IsDir true item Items NumDirs NumFiles
	// false no Items NumDirs NumFiles
	IsDir     bool       `json:"isDir"`
	IsSymlink bool       `json:"isSymlink"`
	Items     []FileInfo `json:"items"`
	NumDirs   int        `json:"numDirs"`
	NumFiles  int        `json:"numFiles"`
	// at file browser path
	Path string `json:"path"`
	// Name file name
	Name string `json:"name"`
	// unit byte
	Size int64 `json:"size"`
	// Extension name of extension
	Extension string `json:"extension"`
	// Modified time utc string
	Modified string `json:"modified"`
	//
	Mode os.FileMode `json:"mode"`
	// Type [ blob text ] text will send content
	Type      string   `json:"type"`
	Subtitles []string `json:"subtitles,omitempty"`
	Content   string   `json:"content,omitempty"`
	// Checksums
	// key will be [ md5 sha1 sha256 sha512 ]
	Checksums map[string]string `json:"checksums,omitempty"`
	Token     string            `json:"token,omitempty"`
}
