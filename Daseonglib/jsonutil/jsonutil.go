package jsonutil

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Game struct {
	ID          string    `json:"id"`
	PackageName string    `json:"packagename"`
	GameTitle   string    `json:"gametitle"`
	GameDesc    *GameDesc `json:"gamedesc"`
}

type GameDesc struct {
	Details1 string `json:"details1"`
	Details2 string `json:"details2"`
}

func GetJsonPath(sPath string) []Game {

	file, err := ioutil.ReadFile(sPath)
	if err != nil {
		return nil
	}

	var Item []Game
	err = json.Unmarshal(file, &Item)

	if err != nil {
		return nil
	}
	//fmt.Println(Item)

	return Item
}

func GetJson(sPath string) bool {

	file, err := ioutil.ReadFile(sPath)
	if err != nil {
		return false
	}

	data := []Game{}
	err = json.Unmarshal(file, &data)

	if err != nil {
		return false
	}
	//fmt.Println(data)

	for _, item := range data {
		sResult := fmt.Sprintf("id:%s packagename:%s gametitle:%s details1:%s details2:%s", item.ID, item.PackageName, item.GameTitle, item.GameDesc.Details1, item.GameDesc.Details2)
		fmt.Println(sResult)
	}
	return true
}

func SetJson(sPath string, data interface{}) bool {

	err := os.MkdirAll(filepath.Dir(sPath), 0755)
	if err != nil {
		return false
	}

	file, err := json.Marshal(data)
	if err != nil {
		return false
	}

	err = ioutil.WriteFile(sPath, file, 0755)
	if err != nil {
		return false
	}
	return true
}

func WriteJsonString(sPath, sText string) {

	file, err := os.OpenFile(sPath, os.O_RDWR|os.O_APPEND, 0660)
	if os.IsNotExist(err) {
		file, err = os.Create(sPath)
	}
	defer file.Close()

	if err != nil {
		return
	}

	n, err := io.WriteString(file, sText)
	if err != nil {
		fmt.Println(n, err)
		return
	}
}
