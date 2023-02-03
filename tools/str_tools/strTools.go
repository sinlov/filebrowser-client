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

func StrArrRemoveDuplicates(slc []string) []string {
	if len(slc) == 0 {
		return slc
	}
	if len(slc) < 1024 {
		return strRemoveDuplicatesByLoop(slc)
	} else {
		return strRemoveDuplicatesByMap(slc)
	}
}

func strRemoveDuplicatesByLoop(slc []string) []string {
	var result []string
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, slc[i])
		}
	}
	return result
}

// strRemoveDuplicatesByMap
// must use slc size gather than 1024
func strRemoveDuplicatesByMap(slc []string) []string {
	var result []string
	tempMap := make(map[string]byte, 1024)
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // If map length changes after map is added, elements do not duplicate
			result = append(result, e)
		}
	}
	return result
}
