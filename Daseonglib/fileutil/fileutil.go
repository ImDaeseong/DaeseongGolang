package fileutil

import (
	_ "fmt"
	"os"
	"path/filepath"
)

//폴더,파일 존재 여부 확인
func IsDirExist(sPath string) bool {
	_, err := os.Stat(sPath)
	return err == nil || os.IsExist(err)
}

func GetFilePath(sFilename string) string {
	return filepath.Dir(sFilename)
}

func GetFileName(sFilename string) string {
	return filepath.Base(sFilename)
}

//폴더 존재 여부 확인
func IsDir(sPath string) bool {
	fileStat, err := os.Stat(sPath)
	if err != nil {
		return false
	}
	return fileStat.IsDir()
}
