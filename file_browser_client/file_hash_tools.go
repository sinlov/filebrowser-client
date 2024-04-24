package file_browser_client

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/sinlov/filebrowser-client/file_browser_log"
	"io"
	"os"
)

func FileSha256Bytes(path string) ([]byte, error) {
	fi, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("want FileSha256Bytes not exist at: %s", path)
	}
	if fi.IsDir() {
		return nil, fmt.Errorf("want FileSha256Bytes is dir at: %s", path)
	}
	openFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("want FileSha256Bytes Open at path: %s, err: %v", path, err)
	}
	defer func(openFile *os.File) {
		errCloseFile := openFile.Close()
		if err != nil {
			file_browser_log.Warnf("FileSha256Bytes close file err: %v", errCloseFile)
		}
	}(openFile)

	buf := make([]byte, 1024*1024)
	h := sha256.New()
	for {
		bytesRead, errFile := openFile.Read(buf)
		if errFile != nil {
			if errFile != io.EOF {
				return nil, fmt.Errorf("want FileSha256Bytes read at path: %s, err: %v", path, errFile)
			}

			break
		}
		h.Write(buf[:bytesRead])
	}

	return h.Sum(nil), nil
}

func FileSha256Hex(path string) (string, error) {
	fileSha256Bytes, err := FileSha256Bytes(path)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(fileSha256Bytes), err
}

func FileMd5Bytes(path string) ([]byte, error) {
	fi, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("want FileSha256Bytes not exist at: %s", path)
	}
	if fi.IsDir() {
		return nil, fmt.Errorf("want FileSha256Bytes is dir at: %s", path)
	}
	openFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("want FileSha256Bytes Open at path: %s, err: %v", path, err)
	}
	defer func(openFile *os.File) {
		errCloseFile := openFile.Close()
		if err != nil {
			file_browser_log.Warnf("FileSha256Bytes close file err: %v", errCloseFile)
		}
	}(openFile)

	buf := make([]byte, 1024*1024)
	h := md5.New()
	for {
		bytesRead, errFile := openFile.Read(buf)
		if errFile != nil {
			if errFile != io.EOF {
				return nil, fmt.Errorf("want FileSha256Bytes read at path: %s, err: %v", path, errFile)
			}

			break
		}
		h.Write(buf[:bytesRead])
	}

	return h.Sum(nil), nil
}

func FileMd5Hex(path string) (string, error) {
	fileSha256Bytes, err := FileMd5Bytes(path)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(fileSha256Bytes), err
}
