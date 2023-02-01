package file_browser_client_test

import (
	"fmt"
	"path"
	"path/filepath"
)

func initTestDataPostFileDir() (string, error) {
	testDataFolderPath, err := getOrCreateTestDataFolderFullPath()
	if err != nil {
		return "", err
	}
	testPostDataFolderPath := path.Join(testDataFolderPath, "post")

	rootLevCnt := 3

	err = addTextFileByTry(testPostDataFolderPath, "data", "json", rootLevCnt)
	if err != nil {
		return "", err
	}

	innerLev1JsonCnt := 5
	innerLev1Folder := filepath.Join(testPostDataFolderPath, "inner_1")
	err = addTextFileByTry(innerLev1Folder, "data", "json", innerLev1JsonCnt)
	if err != nil {
		return "", err
	}

	innerLev11JsonCnt := 4
	innerLev11TxtCnt := 3
	innerLev11Folder := filepath.Join(innerLev1Folder, "inner_1_1")
	err = addTextFileByTry(innerLev11Folder, "data", "json", innerLev11JsonCnt)
	if err != nil {
		return "", err
	}
	err = addTextFileByTry(innerLev11Folder, "log", "txt", innerLev11TxtCnt)
	if err != nil {
		return "", err
	}

	innerLev111JsonCnt := 4
	innerLev111TxtCnt := 3
	innerLev111Folder := filepath.Join(innerLev1Folder, "inner_1_1_1")
	err = addTextFileByTry(innerLev111Folder, "data", "json", innerLev111JsonCnt)
	if err != nil {
		return "", err
	}
	err = addTextFileByTry(innerLev111Folder, "log", "txt", innerLev111TxtCnt)
	if err != nil {
		return "", err
	}

	innerLev12JsonCnt := 4
	innerLev12TxtCnt := 3
	innerLev12Folder := filepath.Join(innerLev1Folder, "inner_1_2")
	err = addTextFileByTry(innerLev12Folder, "data", "json", innerLev12JsonCnt)
	if err != nil {
		return "", err
	}
	err = addTextFileByTry(innerLev12Folder, "log", "txt", innerLev12TxtCnt)
	if err != nil {
		return "", err
	}

	//rootWalkFilesByJson, err := folder.WalkAllFileBySuffix(testDataFolderPath, "json")
	//if err != nil {
	//	return err
	//}

	return testPostDataFolderPath, nil
}

func initTestDataDownloadDir() (string, error) {
	testDataFolderPath, err := getOrCreateTestDataFolderFullPath()
	if err != nil {
		return "", err
	}
	testDownloadDataFolderPath := path.Join(testDataFolderPath, "download")
	if !pathExistsFast(testDownloadDataFolderPath) {
		errMkdir := mkdir(testDownloadDataFolderPath)
		if errMkdir != nil {
			return "", errMkdir
		}
	}
	return testDownloadDataFolderPath, err
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
