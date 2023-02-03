package folder

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// WalkAllByGlob
// can walk all path then return as list, by glob with filepath.Glob
func WalkAllByGlob(path string, glob string, ignoreFolder bool) ([]string, error) {
	fiRoot, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("want Walk not exist at path: %s", path)
		}
		return nil, fmt.Errorf("want Walk not read at path: %s , err: %v", path, err)
	}
	if !fiRoot.IsDir() {
		return nil, fmt.Errorf("want Walk path is file, at: %s", path)
	}
	pathOfGlob := fmt.Sprintf("%s%s%s", path, string(filepath.Separator), glob)
	matches, err := filepath.Glob(pathOfGlob)
	if err != nil {
		return nil, fmt.Errorf("want Walk by path %s by glob %s ,err: %v", path, glob, err)
	}
	if len(matches) == 0 {
		return nil, nil
	}
	files := make([]string, 0, 30)
	for _, match := range matches {
		f, errStat := os.Stat(match)
		if errStat != nil {
			return nil, fmt.Errorf("want Walk Stat at path %s by glob %s ,err: %v", match, glob, errStat)
		}
		if ignoreFolder && f.IsDir() {
			continue
		}
		files = append(files, match)
	}

	return files, nil
}

// WalkAllByMatchPath
// can walk all path then return as list, by under path pattern
func WalkAllByMatchPath(path string, pattern string, ignoreFolder bool) ([]string, error) {
	fiRoot, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("want Walk not exist at path: %s", path)
		}
		return nil, fmt.Errorf("want Walk not read at path: %s , err: %v", path, err)
	}
	if !fiRoot.IsDir() {
		return nil, fmt.Errorf("want Walk path is file, at: %s", path)
	}
	files := make([]string, 0, 30)
	err = filepath.Walk(path, func(filename string, fi os.FileInfo, err error) error {
		if ignoreFolder && fi.IsDir() { // ignore dir
			return nil
		}
		innerPath := strings.Replace(filename, path, "", -1)
		innerPath = strings.TrimPrefix(innerPath, string(filepath.Separator))
		matched, err := regexp.MatchString(pattern, innerPath)
		if err != nil {
			return nil
		}
		if matched {
			files = append(files, filename)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("now Walk path: %s , err: %v", path, err)
	}

	return files, nil
}

// WalkAllFileBySuffix
// can walk all file then return all files as list, by file suffix
func WalkAllFileBySuffix(path string, suffix string) ([]string, error) {
	fiRoot, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("want Walk not exist at path: %s", path)
		}
		return nil, fmt.Errorf("want Walk not read at path: %s , err: %v", path, err)
	}
	if !fiRoot.IsDir() {
		return nil, fmt.Errorf("want Walk path is file, at: %s", path)
	}
	files := make([]string, 0, 30)
	err = filepath.Walk(path, func(filename string, fi os.FileInfo, err error) error {
		if fi.IsDir() { // ignore dir
			return nil
		}
		if strings.HasSuffix(fi.Name(), suffix) {
			files = append(files, filename)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("now Walk path: %s , err: %v", path, err)
	}

	return files, nil
}
