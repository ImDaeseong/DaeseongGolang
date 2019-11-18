package iniutil

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type iniTag struct {
	sSection  string
	sKeyname  string
	sKeyValue string
}

var (
	iMap     map[string]iniTag
	isecMap  map[string]string
	Index    int = 0
	secIndex int = 0
)

func GetiniList() {

	for key, _ := range iMap {
		sResult := fmt.Sprintf("%s %s %s %s", key, iMap[key].sSection, iMap[key].sKeyname, iMap[key].sKeyValue)
		fmt.Println(sResult)
	}
}

func Removeini(sSection string) {

	var sSect string

	//[] 제거
	Replacer := strings.NewReplacer(
		"[", "",
		"]", "",
	)

	for key, _ := range iMap {

		sSect = Replacer.Replace(iMap[key].sSection)
		if sSect == sSection {
			delete(iMap, key)
		}
	}
}

func Addini(skey, sSection, sKeyname, sKeyValue string) {
	iMap[skey] = iniTag{sSection, sKeyname, sKeyValue}
}

func sortmap() {

	//key sort
	ikeys := make([]string, 0, len(isecMap))
	for ik := range isecMap {
		ikeys = append(ikeys, ik)
	}
	sort.Strings(ikeys)

	//data sort
	keys := make([]string, 0, len(iMap))
	for k := range iMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
}

func writeText(file *os.File, sText string) bool {

	n, err := io.WriteString(file, sText)
	if err != nil {
		fmt.Println(n, err)
		return false
	}
	return true
}

func Getloadini(sPath string) bool {

	iMap = make(map[string]iniTag)
	isecMap = make(map[string]string)

	f, err := os.OpenFile(sPath, os.O_RDONLY, 0644)
	if err != nil {
		return false
	}

	var sTag string

	reader := bufio.NewScanner(f)
	for reader.Scan() {
		line := reader.Text()
		line = strings.TrimSpace(line)

		if line[0] == '[' && line[len(line)-1] == ']' {

			sTag = line

			skey := fmt.Sprintf("%d", secIndex)

			isecMap[skey] = sTag

			secIndex++

		} else if strings.Contains(line, "=") {

			parts := strings.SplitN(line, "=", 2)
			//fmt.Printf("%s %s %s\n", sTag, parts[0], parts[1])

			skey := fmt.Sprintf("%d", Index)

			Addini(skey, sTag, parts[0], parts[1])

			Index++
		}
	}

	return true
}

func Setloadini(sPath string) bool {

	file, err := os.OpenFile(sPath, os.O_RDWR|os.O_APPEND, 0660)
	if os.IsNotExist(err) {
		file, err = os.Create(sPath)
	}
	defer file.Close()

	if err != nil {
		return false
	}

	//sort
	sortmap()

	for ikey, _ := range isecMap {

		//key write
		sResult := fmt.Sprintf("%s\r\n", isecMap[ikey])
		writeText(file, sResult)

		//data write
		for key, _ := range iMap {

			if isecMap[ikey] == iMap[key].sSection {

				sResult := fmt.Sprintf("%s=%s\r\n", iMap[key].sKeyname, iMap[key].sKeyValue)
				writeText(file, sResult)
			}
		}

		//key end
		sResult = fmt.Sprintf("\r\n")
		writeText(file, sResult)
	}

	return true
}

func GetProfileString(sSection string, sKeyname string) string {

	var sResult string = ""
	var sSect string

	//[] 제거
	Replacer := strings.NewReplacer(
		"[", "",
		"]", "",
	)

	for key, _ := range iMap {

		sSect = Replacer.Replace(iMap[key].sSection)
		if sSect == sSection && sKeyname == iMap[key].sKeyname {
			sResult = iMap[key].sKeyValue
			break
		}
	}
	return sResult
}

func SetProfileString(sSection string, sKeyname string, sKeyValue string) {

	var sResultkey string = ""
	var sSect string

	//[] 제거
	Replacer := strings.NewReplacer(
		"[", "",
		"]", "",
	)

	for key, _ := range iMap {

		sSect = Replacer.Replace(iMap[key].sSection)
		if sSect == sSection && sKeyname == iMap[key].sKeyname {
			sResultkey = key
			break
		}
	}

	if sResultkey != "" {

		iMap[sResultkey] = iniTag{"[" + sSection + "]", sKeyname, sKeyValue}

	} else {

		skey := fmt.Sprintf("%d", secIndex)
		isecMap[skey] = "[" + sSection + "]"
		secIndex++

		skey = fmt.Sprintf("%d", Index)
		iMap[skey] = iniTag{"[" + sSection + "]", sKeyname, sKeyValue}
		Index++
	}
}
