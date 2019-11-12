// stringutil
package stringutil

import (
	_ "fmt"
	"strings"
)

//파일 확장자
func GetExt(filename string) string {

	dotIndex := strings.LastIndex(filename, ".")
	if dotIndex == -1 {
		return ""
	}

	return string(filename[dotIndex+1:])
}

//파일 이름
func GetFileName(filename string) string {

	slashIndex := strings.LastIndex(filename, "\\")
	if slashIndex == -1 {
		return ""
	}

	return string(filename[slashIndex+1:])
}

//only 파일 이름
func GetOnlyFileName(filename string) string {

	slashIndex := strings.LastIndex(filename, "\\")
	if slashIndex == -1 {
		return ""
	}

	dotIndex := strings.LastIndex(filename, ".")
	if dotIndex == -1 {
		return ""
	}

	//fmt.Println(slashIndex)
	//fmt.Println(dotIndex)

	return string(filename[slashIndex+1 : dotIndex])
}
