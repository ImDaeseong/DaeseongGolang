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
	iMap     map[int]iniTag
	isecMap  map[int]string
	Index    int = 0
	secIndex int = 0
)

func GetiniList() {

	for key, _ := range iMap {
		sResult := fmt.Sprintf("%d %s %s %s", key, iMap[key].sSection, iMap[key].sKeyname, iMap[key].sKeyValue)
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

func Addini(nkey int, sSection, sKeyname, sKeyValue string) {
	iMap[nkey] = iniTag{sSection, sKeyname, sKeyValue}
}

func isSection(sSection string) bool {

	var bSect bool = false
	for key, _ := range isecMap {

		if isecMap[key] == sSection {
			bSect = true
			break
		}
	}
	return bSect

}

func sort_isecMap(m map[int]string) []int {
	keys := []int{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Sort(sort.IntSlice(keys))
	return keys
}

func sort_iMap(m map[int]iniTag) []int {
	keys := []int{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Sort(sort.IntSlice(keys))
	return keys
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

	iMap = make(map[int]iniTag)
	isecMap = make(map[int]string)

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
			isecMap[secIndex] = sTag
			secIndex++

		} else if strings.Contains(line, "=") {

			parts := strings.SplitN(line, "=", 2)
			Addini(Index, sTag, parts[0], parts[1])
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

	for ikey, _ := range sort_isecMap(isecMap) {

		//key write
		sResult := fmt.Sprintf("%s\r\n", isecMap[ikey])
		writeText(file, sResult)

		//data write
		for key, _ := range sort_iMap(iMap) {

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

	var bUpdate bool = false
	var nkey int = 0
	var sSect string

	//[] 제거
	Replacer := strings.NewReplacer(
		"[", "",
		"]", "",
	)

	for key, _ := range iMap {

		sSect = Replacer.Replace(iMap[key].sSection)
		if sSect == sSection && sKeyname == iMap[key].sKeyname {
			nkey = key
			bUpdate = true
			break
		}
	}

	if bUpdate == true {

		sSect := fmt.Sprintf("[%s]", sSection)
		iMap[nkey] = iniTag{sSect, sKeyname, sKeyValue}

	} else {

		sSect := fmt.Sprintf("[%s]", sSection)

		if isSection(sSect) == false {
			isecMap[secIndex] = sSect
			secIndex++
		}

		iMap[Index] = iniTag{sSect, sKeyname, sKeyValue}
		Index++
	}
}
