package iniutil

import (
	"bufio"
	"fmt"
	"os"
	_ "sort"
	"strings"
)

type iniTag struct {
	sSection  string
	sKeyname  string
	sKeyValue string
}

var (
	iniList = []iniTag{}
)

func Getloadini(sPath string) bool {

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

		} else if strings.Contains(line, "=") {

			parts := strings.SplitN(line, "=", 2)
			//fmt.Printf("%s %s %s\n", sTag, parts[0], parts[1])

			ini := iniTag{}
			ini.sSection = sTag
			ini.sKeyname = parts[0]
			ini.sKeyValue = parts[1]
			iniList = append(iniList, ini)
		}
	}

	/*
		for _, item := range iniList {
			sResult := fmt.Sprintf("%s %s %s", item.sSection, item.sKeyname, item.sKeyValue)
			fmt.Println(sResult)
		}
	*/

	return true
}

func GetProfileString(sSection string, sKeyname string) string {

	var sResult string = ""
	var sSect string = ""

	Replacer := strings.NewReplacer(
		"[", "",
		"]", "",
	)

	for _, item := range iniList {

		//[] 제거
		sSect = Replacer.Replace(item.sSection)

		if sSect == sSection && item.sKeyname == sKeyname {

			sResult = item.sKeyValue
			break
		}
	}

	return sResult
}

func SetProfileString(sSection string, sKeyname string, sKeyValue string) bool {

	sTag := fmt.Sprintf("[%s]", sSection)

	ini := iniTag{}
	ini.sSection = sTag
	ini.sKeyname = sKeyname
	ini.sKeyValue = sKeyValue
	iniList = append(iniList, ini)

	for _, item := range iniList {
		sResult := fmt.Sprintf("%s %s %s", item.sSection, item.sKeyname, item.sKeyValue)
		fmt.Println(sResult)
	}

	return true
}
