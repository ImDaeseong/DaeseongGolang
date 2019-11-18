package downloadutil

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetWebPage(sUrl string) (string, error) {
	req, err := http.NewRequest("GET", sUrl, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("User-Agent", "Daeseonglib")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	byte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(byte), nil
}

func GetDownloadFile(sUrl, sSavePath string) string {

	bByte, err := downloadfile(sUrl)
	if err != nil {
		return ""
	}

	filename := getfileName(sUrl)
	filepath := fmt.Sprintf("%s\\%s", sSavePath, filename)

	bWrite := writefile(filepath, bByte)
	if bWrite {
		return filepath
	}

	return ""
}

func getfileName(filename string) string {

	slashIndex := strings.LastIndex(filename, "/")
	if slashIndex == -1 {
		return ""
	}

	return string(filename[slashIndex+1:])
}

func downloadfile(sUrl string) ([]byte, error) {
	req, err := http.NewRequest("GET", sUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "downloadutil")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bByte, nil
}

func writefile(sFileName string, bByte []byte) bool {
	err := ioutil.WriteFile(sFileName, bByte, 0644)
	if err != nil {
		return false
	}
	return true
}
