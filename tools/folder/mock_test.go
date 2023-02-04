package folder

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

// test case file tools start

var currentTestDataFolderAbsPath = ""

func getOrCreateTestDataFolderFullPath() (string, error) {
	if currentTestDataFolderAbsPath != "" {
		return currentTestDataFolderAbsPath, nil
	}
	currentFolderPath, err := getCurrentFolderPath()
	if err != nil {
		return "", err
	}
	currentTestDataFolderAbsPath = filepath.Join(currentFolderPath, "testdata")
	if !pathExistsFast(currentTestDataFolderAbsPath) {
		err := mkdir(currentTestDataFolderAbsPath)
		if err != nil {
			currentTestDataFolderAbsPath = ""
			return "", err
		}
	}
	return currentTestDataFolderAbsPath, nil
}

func goldenDataSaveFast(t *testing.T, data []byte, extraName string) error {
	return goldenDataSave(t, data, extraName, os.FileMode(0766))
}

func goldenDataSave(t *testing.T, data []byte, extraName string, fileMod fs.FileMode) error {
	testDataFolderFullPath, err := getOrCreateTestDataFolderFullPath()
	if err != nil {
		return fmt.Errorf("try goldenDataSave err: %v", err)
	}
	savePath := filepath.Join(testDataFolderFullPath, fmt.Sprintf("%s-%s.golden", t.Name(), extraName))
	err = writeFileByByte(savePath, data, fileMod, true)
	if err != nil {
		return fmt.Errorf("try goldenDataSave at path: %s err: %v", savePath, err)
	}
	return nil
}

func goldenDataReadAsByte(t *testing.T, extraName string) ([]byte, error) {
	testDataFolderFullPath, err := getOrCreateTestDataFolderFullPath()
	if err != nil {
		return nil, fmt.Errorf("try goldenDataReadAsByte err: %v", err)
	}

	savePath := filepath.Join(testDataFolderFullPath, fmt.Sprintf("%s-%s.golden", t.Name(), extraName))

	fileAsByte, err := readFileAsByte(savePath)
	if err != nil {
		return nil, fmt.Errorf("try goldenDataReadAsByte err: %v", err)
	}
	return fileAsByte, nil
}

// getCurrentFolderPath can get run path this golang dir
func getCurrentFolderPath() (string, error) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("can not get current file info")
	}
	return filepath.Dir(file), nil
}

// pathExists path exists
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// pathExistsFast path exists fast
func pathExistsFast(path string) bool {
	exists, _ := pathExists(path)
	return exists
}

// rmDir remove dir by path
func rmDir(path string, force bool) error {
	if force {
		return os.RemoveAll(path)
	}
	exists, err := pathExists(path)
	if err != nil {
		return err
	}
	if exists {
		return os.RemoveAll(path)
	}
	return fmt.Errorf("remove dir not exist path: %s , use force can cover this err", path)
}

// mkdir will use FileMode 0766
func mkdir(path string) error {
	err := os.MkdirAll(path, os.FileMode(0766))
	if err != nil {
		return fmt.Errorf("fail MkdirAll at path: %s , err: %v", path, err)
	}
	return nil
}

// readFileAsByte read file as byte array
func readFileAsByte(path string) ([]byte, error) {
	exists, err := pathExists(path)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("path not exist %v", path)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("path: %s read err: %v", path, err)
	}

	return data, nil
}

// readFileAsJson read file as json
func readFileAsJson(path string, v interface{}) error {
	fileAsByte, err := readFileAsByte(path)
	err = json.Unmarshal(fileAsByte, v)
	if err != nil {
		return fmt.Errorf("path: %s , read file as json err: %v", path, err)
	}
	return nil
}

// writeFileByByte write bytes to file
// path most use Abs Path
// data []byte
// fileMod os.FileMode(0666) or os.FileMode(0644)
// coverage true will coverage old
func writeFileByByte(path string, data []byte, fileMod fs.FileMode, coverage bool) error {
	if !coverage {
		exists, err := pathExists(path)
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("not coverage, which path exist %v", path)
		}
	}
	parentPath := filepath.Dir(path)
	if !pathExistsFast(parentPath) {
		err := os.MkdirAll(parentPath, fileMod)
		if err != nil {
			return fmt.Errorf("can not WriteFileByByte at new dir at mode: %v , at parent path: %v", fileMod, parentPath)
		}
	}
	err := os.WriteFile(path, data, fileMod)
	if err != nil {
		return fmt.Errorf("write data at path: %v, err: %v", path, err)
	}
	return nil
}

// writeFileAsJson write json file
// path most use Abs Path
// v data
// fileMod os.FileMode(0666) or os.FileMode(0644)
// coverage true will coverage old
// beauty will format json when write
func writeFileAsJson(path string, v interface{}, fileMod fs.FileMode, coverage, beauty bool) error {
	if !coverage {
		exists, err := pathExists(path)
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("not coverage, which path exist %v", path)
		}
	}
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	if beauty {
		var str bytes.Buffer
		err := json.Indent(&str, data, "", "  ")
		if err != nil {
			return err
		}
		return writeFileByByte(path, str.Bytes(), fileMod, coverage)
	}
	return writeFileByByte(path, data, fileMod, coverage)
}

// writeFileAsJsonBeauty write json file as 0766 and beauty
func writeFileAsJsonBeauty(path string, v interface{}, coverage bool) error {
	return writeFileAsJson(path, v, os.FileMode(0766), coverage, true)
}

func addTextFileByTry(targetDir, fileHead, suffix string, cnt int) error {

	if !pathExistsFast(targetDir) {
		err := mkdir(targetDir)
		if err != nil {
			return err
		}
	}

	var foo struct {
		Foo int    `json:"foo"`
		Bar string `json:"bar"`
	}

	for i := 0; i < cnt; i++ {
		fName := fmt.Sprintf("%s_%d.%s", fileHead, i, suffix)
		newJsonPath := filepath.Join(targetDir, fName)
		foo.Foo = i
		errJsonWrite := writeFileAsJsonBeauty(newJsonPath, foo, true)
		if errJsonWrite != nil {
			return errJsonWrite
		}
	}
	return nil
}

// randomStr
// new random string by cnt
func randomStr(cnt uint) string {
	var letters = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	result := make([]byte, cnt)
	keyL := len(letters)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(keyL)]
	}
	return string(result)
}

// randomInt
// new random int by max
func randomInt(max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max)
}

// test case file tools end
