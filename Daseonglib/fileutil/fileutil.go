package fileutil

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
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

func DeleteFile(sPath string) error {
	file, err := os.Stat(sPath)
	if err != nil {
		return err
	}

	if file.IsDir() {
		err := os.RemoveAll(sPath)
		if err != nil {
			return err
		}
	} else {
		err := os.Remove(sPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadLineFile(sPath string) {
	var line int = 0

	file, err := os.Open(sPath)
	if err != nil {
		return
	}

	fileScan := bufio.NewScanner(file)
	for fileScan.Scan() {
		line++
		fmt.Printf("[%d:%s]\n", line, fileScan.Text())
	}
}

func FindFileList(sPath string) []string {
	fileList := []string{}
	filepath.Walk(sPath, func(p string, f os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !f.IsDir() {
			fileList = append(fileList, p)
		}
		return nil
	})
	sort.Strings(fileList)
	return fileList
}

func FindDirList(sPath string) []string {
	fileList := []string{}

	filepath.Walk(sPath, func(p string, f os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if f.IsDir() {
			fileList = append(fileList, p)
		}
		return nil
	})
	return fileList
}
