package dictutil

import (
	"fmt"
)

var (
	words map[string]string
)

func InitWords() {
	words = make(map[string]string)
}

func GetWordsList() {

	for key, value := range words {
		fmt.Println(key + " = " + value)
	}
}

func IsDictionary(skey string) bool {
	_, exists := words[skey]
	return exists
}

func RemoveWords(skey string) bool {

	if IsDictionary(skey) {
		delete(words, skey)
		return true
	}
	return false
}

func AddWords(skey, sValue string) bool {

	if IsDictionary(skey) {
		return false
	}
	words[skey] = sValue
	return true
}

func UpdateWords(skey, sValue string) bool {

	if IsDictionary(skey) {
		words[skey] = sValue
		return true
	}
	return false
}
