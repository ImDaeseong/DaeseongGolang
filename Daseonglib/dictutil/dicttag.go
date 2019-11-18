package dictutil

import (
	"fmt"
)

type DicTag struct {
	sValue1 string
	sValue2 string
	sValue3 string
}

var (
	dicTags map[string]DicTag
)

func InitTags() {
	dicTags = make(map[string]DicTag)
}

func GetTagsList() {

	for key, value := range dicTags {
		sResult := fmt.Sprintf("key=%s, value=%s", key, value)
		fmt.Println(sResult)
	}

	for key, _ := range dicTags {
		sResult := fmt.Sprintf("key=%s sValue1=%s sValue2=%s sValue3=%s", key, dicTags[key].sValue1, dicTags[key].sValue2, dicTags[key].sValue3)
		fmt.Println(sResult)
	}
}

func IsTags(skey string) bool {
	_, exists := dicTags[skey]
	return exists
}

func RemoveTags(skey string) bool {

	if IsTags(skey) {
		delete(dicTags, skey)
		return true
	}
	return false
}

func AddTags(skey, sValue1, sValue2, sValue3 string) bool {

	if IsTags(skey) {
		return false
	}
	dicTags[skey] = DicTag{sValue1, sValue2, sValue3}
	return true
}

func UpdateTags(skey, sValue1, sValue2, sValue3 string) bool {

	if IsTags(skey) {
		dicTags[skey] = DicTag{sValue1, sValue2, sValue3}
		return true
	}
	return false
}
