package tools

import (
	"sort"
	"strings"
)

func StrInArr(target string, strArray []string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}

func Str2LineRaw(target string) string {
	newStr := strings.Replace(target, "\r\n", `\n`, -1)
	newStr = strings.Replace(newStr, "\n", `\n`, -1)
	newStr = strings.Replace(newStr, "\r", `\n`, -1)
	return newStr
}
