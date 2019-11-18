// main
package main

import (
	"Daseonglib/dictutil"
	"Daseonglib/fileutil"
	"Daseonglib/iniutil"
	"Daseonglib/iniwinutil"
	"Daseonglib/jsonutil"
	"Daseonglib/stringutil"
	"Daseonglib/timeutil"
	"fmt"
	_ "registryutil" //internal 파일 경로는 Go\src\registryutil
	"strconv"
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

func getjsonutil() {

	var gamedata []Game

	//json file create
	gamedata = append(gamedata, Game{ID: "1", PackageName: "com.pearlabyss.blackdesertm", GameTitle: "검은사막 모바일", GameDesc: &GameDesc{Details1: "당신이 진짜로 원했던 모험의 시작", Details2: "월드클래스 MMORPG “검은사막 모바일”"}})
	gamedata = append(gamedata, Game{ID: "2", PackageName: "com.kakaogames.moonlight", GameTitle: "달빛조각사", GameDesc: &GameDesc{Details1: "500만 구독자의 게임 판타지 대작 '달빛조각사'", Details2: "- 5레벨만 달성해도 달빛조각사 이모티콘 100% 지급!"}})
	gamedata = append(gamedata, Game{ID: "3", PackageName: "com.ncsoft.lineagem19", GameTitle: "리니지M", GameDesc: &GameDesc{Details1: "PC의 향수! 리니지 본질 그대로 리니지M", Details2: "PC리니지와 동일한 아덴월드의 오픈 필드"}})
	gamedata = append(gamedata, Game{ID: "4", PackageName: "com.netmarble.bnsmkr", GameTitle: "블레이드&소울 레볼루션", GameDesc: &GameDesc{Details1: "원작 감성의 방대한 세계관과 복수 중심의 흥미진진한 스토리", Details2: "MMORPG의 필드를 제대로 즐길 수 있는 경공"}})
	gamedata = append(gamedata, Game{ID: "5", PackageName: "com.cjenm.sknights", GameTitle: "세븐나이츠", GameDesc: &GameDesc{Details1: "Netmarble롤플레잉", Details2: "세나의 재탄생, 세븐나이츠: 리부트"}})
	gamedata = append(gamedata, Game{ID: "6", PackageName: "com.google.android.youtube", GameTitle: "YouTube", GameDesc: &GameDesc{Details1: "Google LLC동영상 플레이어/편집기", Details2: "좋아하는 동영상 빠르게 검색하기"}})
	jsonutil.SetJson("c:\\game.json", gamedata)

	//json file read
	jsonutil.GetJson("c:\\game.json")

	/*
		//json file read
		item := jsonutil.GetJsonPath("c:\\game.json")
		for i := range item {
			sResult := fmt.Sprintf("id:%s packagename:%s gametitle:%s details1:%s details2:%s", item[i].ID, item[i].PackageName, item[i].GameTitle, item[i].GameDesc.Details1, item[i].GameDesc.Details2)
			fmt.Println(sResult)
		}
	*/

	/*
		var sVal string
		sVal = fmt.Sprintf("[{\"ID\": \"%s\", \"PackageName\": \"%s\", \"GameTitle\": \"%s\", \"GameDesc\": {\"Details1\": \"%s\", \"Details2\": \"%s\"} },",
			"1", "com.pearlabyss.blackdesertm", "검은사막 모바일", "당신이 진짜로 원했던 모험의 시작", "월드클래스 MMORPG “검은사막 모바일")
		jsonutil.WriteJsonString("c:\\game.json", sVal)

		sVal = fmt.Sprintf("{\"ID\": \"%s\", \"PackageName\": \"%s\", \"GameTitle\": \"%s\", \"GameDesc\": {\"Details1\": \"%s\", \"Details2\": \"%s\"} },",
			"2", "com.kakaogames.moonlight", "달빛조각사", "500만 구독자의 게임 판타지 대작 '달빛조각사'", "- 5레벨만 달성해도 달빛조각사 이모티콘 100% 지급!")
		jsonutil.WriteJsonString("c:\\game.json", sVal)

		sVal = fmt.Sprintf("{\"ID\": \"%s\", \"PackageName\": \"%s\", \"GameTitle\": \"%s\", \"GameDesc\": {\"Details1\": \"%s\", \"Details2\": \"%s\"} },",
			"3", "com.ncsoft.lineagem19", "리니지M", "PC의 향수! 리니지 본질 그대로 리니지M", "PC리니지와 동일한 아덴월드의 오픈 필드")
		jsonutil.WriteJsonString("c:\\game.json", sVal)

		sVal = fmt.Sprintf("{\"ID\": \"%s\", \"PackageName\": \"%s\", \"GameTitle\": \"%s\", \"GameDesc\": {\"Details1\": \"%s\", \"Details2\": \"%s\"} },",
			"4", "com.netmarble.bnsmkr", "블레이드&소울 레볼루션", "원작 감성의 방대한 세계관과 복수 중심의 흥미진진한 스토리", "MMORPG의 필드를 제대로 즐길 수 있는 경공")
		jsonutil.WriteJsonString("c:\\game.json", sVal)

		sVal = fmt.Sprintf("{\"ID\": \"%s\", \"PackageName\": \"%s\", \"GameTitle\": \"%s\", \"GameDesc\": {\"Details1\": \"%s\", \"Details2\": \"%s\"} },",
			"5", "com.cjenm.sknights", "세븐나이츠", "Netmarble롤플레잉", "세나의 재탄생, 세븐나이츠: 리부트")
		jsonutil.WriteJsonString("c:\\game.json", sVal)

		sVal = fmt.Sprintf("{\"ID\": \"%s\", \"PackageName\": \"%s\", \"GameTitle\": \"%s\", \"GameDesc\": {\"Details1\": \"%s\", \"Details2\": \"%s\"} }]",
			"6", "com.google.android.youtube", "YouTube", "Google LLC동영상 플레이어/편집기", "좋아하는 동영상 빠르게 검색하기")
		jsonutil.WriteJsonString("c:\\game.json", sVal)
	*/

	/*
		//json file read
		item := jsonutil.GetJsonPath("c:\\game.json")
		for i := range item {
			sResult := fmt.Sprintf("id:%s packagename:%s gametitle:%s details1:%s details2:%s", item[i].ID, item[i].PackageName, item[i].GameTitle, item[i].GameDesc.Details1, item[i].GameDesc.Details2)
			fmt.Println(sResult)
		}
	*/
}

func getregistryutil() {

	/*
		if registryutil.SetregistryString("Software\\Daeseong\\Daeseong", "GameList", "3") {

			nCount, _ := registryutil.GetregistryString("Software\\Daeseong\\Daeseong", "GameList")
			fmt.Println(nCount)

		}

		if registryutil.SetregistryDWord("Software\\Daeseong\\Daeseong", "Gamekey1", 1) {

			nValue, _ := registryutil.GetregistryDWord("Software\\Daeseong\\Daeseong", "Gamekey1")
			fmt.Println(nValue)
		}
	*/

	/*
		if registryutil.SetregistryStringWOW64("Software\\Daeseong\\Daeseong", "GameList", "3") {

			nCount, _ := registryutil.GetregistryStringWOW64("Software\\Daeseong\\Daeseong", "GameList")
			fmt.Println(nCount)

		}

		if registryutil.SetregistryDWordWOW64("Software\\Daeseong\\Daeseong", "Gamekey1", 1) {

			nValue, _ := registryutil.GetregistryDWordWOW64("Software\\Daeseong\\Daeseong", "Gamekey1")
			fmt.Println(nValue)
		}
	*/

	/*
		if registryutil.SetregistryString32("Software\\Daeseong\\Daeseong", "GameList", "3") {

			nCount, _ := registryutil.GetregistryString32("Software\\Daeseong\\Daeseong", "GameList")
			fmt.Println(nCount)

		}

		if registryutil.SetregistryDWord32("Software\\Daeseong\\Daeseong", "Gamekey1", 1) {

			nValue, _ := registryutil.GetregistryDWord32("Software\\Daeseong\\Daeseong", "Gamekey1")
			fmt.Println(nValue)
		}
	*/

	/*
		err := registryutil.DeleteValueregistry32("Software\\Daeseong\\Daeseong", "GameList")

		err = registryutil.DeleteKeyregistry32("Software\\Daeseong\\Daeseong")
		if err == nil {

			err = registryutil.DeleteKeyregistry32("Software\\Daeseong\\")
			if err == nil {
				fmt.Println("delete all")
			}
		}
	*/
}

func getiniwinutil() {

	sPath := "c:\\gameinfo.ini"

	nGameCount := 3
	if iniwinutil.SetProfileString("GameList", "GameCount", strconv.Itoa(nGameCount), sPath) {

		for i := 0; i < nGameCount; i++ {
			key := fmt.Sprintf("Gamekey%d", i)
			value := fmt.Sprintf("Gamename%d", i)
			iniwinutil.SetProfileString("GameList", key, value, sPath)
		}

	}

	nReadCount := iniwinutil.GetProfileString("GameList", "GameCount", "", sPath)
	nCount, err := strconv.Atoi(nReadCount)
	if err == nil {

		for i := 0; i < nCount; i++ {

			Readkey := fmt.Sprintf("Gamekey%d", i)
			Readvalue := iniwinutil.GetProfileString("GameList", Readkey, "", sPath)

			key := fmt.Sprintf("GameItem")
			value := fmt.Sprintf("com.kakaogames.moonlight%d", i)
			iniwinutil.SetProfileString(Readvalue, key, value, sPath)
		}
	}

	//ReadGameitem := iniwinutil.GetProfileString("Gamename0", "GameItem", "", sPath)
	//fmt.Println(ReadGameitem)
}

func getstringutil() {

	sResult1 := stringutil.GetExt("E:\\GoApp\\src\\src.zip")
	sResult2 := stringutil.GetFileName("E:\\GoApp\\src\\src.zip")
	sResult3 := stringutil.GetOnlyFileName("E:\\GoApp\\src\\src.zip")
	fmt.Println("확장자: " + sResult1)
	fmt.Println("파일: " + sResult2)
	fmt.Println("파일이름만: " + sResult3)
}

func gettimeutil() {

	fmt.Println(timeutil.GetFullCurrentDay())
	fmt.Println(timeutil.GetCurrentDay())
	fmt.Println(timeutil.GetCurrentTime())
	fmt.Println(timeutil.GetToday())

	fmt.Println("년:" + timeutil.GetYear())
	fmt.Println("월:" + timeutil.GetMonth())
	fmt.Println("일:" + timeutil.GetDay())

	fmt.Println("년(+):" + timeutil.SetYear(1))
	fmt.Println("년(-):" + timeutil.SetYear(-1))

	fmt.Println("월(+):" + timeutil.SetMonth(1))
	fmt.Println("월(-):" + timeutil.SetMonth(-1))

	fmt.Println("일(+):" + timeutil.SetDay(1))
	fmt.Println("일(-):" + timeutil.SetDay(-1))

	//sStart, _ := timeutil.ConvertStrToTime("2019-11-11 12:23:00")
	//sEnd, _ := timeutil.ConvertStrToTime("2019-11-12 12:23:00")
	//diffSeconds := timeutil.SubtimeTime(sStart, sEnd)

	diffSeconds := timeutil.SubstringTime("2019-11-11 12:23:00", "2019-11-12 12:23:00")

	minutes := int(diffSeconds / 60)
	hour := int(minutes / 60)

	fmt.Println("초를 분으로 :" + fmt.Sprintf("%d", minutes))
	fmt.Println("분를 시으로 :" + fmt.Sprintf("%d", hour))
}

func getdict() {

	dictutil.InitWords()

	dictutil.AddWords("lipas", "바퀴벌레")
	dictutil.AddWords("muram", "우울한우울한")
	dictutil.AddWords("persaingan", "경쟁")
	dictutil.AddWords("rawit", "작은")

	if dictutil.IsDictionary("lipas") {
		fmt.Println("is exist")
	}

	dictutil.RemoveWords("lipas")

	dictutil.AddWords("tikung", "커브길")

	dictutil.UpdateWords("muram", "우울한")

	dictutil.GetWordsList()
}

func getdicttag() {

	dictutil.InitTags()

	dictutil.AddTags("lipas", "cockroach", "lipas", "바퀴벌레")
	dictutil.AddTags("muram", "gloomy1", "muram1", "우울한우울한")
	dictutil.AddTags("persaingan", "competition", "persaingan", "경쟁")
	dictutil.AddTags("rawit", "small", "rawit", "작은")

	if dictutil.IsTags("lipas") {
		fmt.Println("is exist")
	}

	dictutil.RemoveTags("lipas")

	dictutil.AddTags("sembelit", "constipation", "sembelit", "변비")

	dictutil.UpdateTags("muram", "gloomy", "muram", "우울한")

	dictutil.GetTagsList()
}

func getfileutil() {

	if fileutil.IsDirExist("D:\\DaeseongGolang\\src\\Daseonglib\\fileutil") {
		fmt.Println("is exist")
	}

	if fileutil.IsDirExist("D:\\DaeseongGolang\\src\\Daseonglib\\fileutil\\fileutil.go") {
		fmt.Println("is exist")
	}

	sfilepath := fileutil.GetFilePath("D:\\DaeseongGolang\\src\\Daseonglib\\fileutil\\fileutil.go")
	fmt.Println(sfilepath)

	sfilename := fileutil.GetFileName("D:\\DaeseongGolang\\src\\Daseonglib\\fileutil\\fileutil.go")
	fmt.Println(sfilename)

	if fileutil.IsDir("D:\\DaeseongGolang\\src\\Daseonglib\\fileutil") {
		fmt.Println("is exist")
	}

}

func getiniutil() {

	/*
		//------gameinfo.ini 내용
		[GameList]
		GameCount=3
		Gamekey0=Gamename0
		Gamekey1=Gamename1
		Gamekey2=Gamename2
		[Gamename0]
		GameItem=com.kakaogames.moonlight0
		[Gamename1]
		GameItem=com.kakaogames.moonlight1
		[Gamename2]
		GameItem=com.kakaogames.moonlight2
	*/

	iniutil.Getloadini("c:\\gameinfo.ini")

	nReadCount := iniutil.GetProfileString("GameList", "GameCount")
	nCount, err := strconv.Atoi(nReadCount)
	if err == nil {

		for i := 0; i < nCount; i++ {

			Gamekey := fmt.Sprintf("Gamekey%d", i)
			Gamename := iniutil.GetProfileString("GameList", Gamekey)
			GameItem := iniutil.GetProfileString(Gamename, "GameItem")

			fmt.Println("[" + Gamekey + "]" + Gamename + " - " + GameItem)
		}
	}

	/*
		iniutil.Removeini("GameList")
		iniutil.GetiniList()
	*/

	iniutil.SetProfileString("GameList", "GameCount", "4")
	iniutil.SetProfileString("GameList", "Gamekey3", "Gamename3")
	iniutil.SetProfileString("Gamename3", "GameItem", "com.kakaogames.moonlight3")
	iniutil.GetiniList()
	iniutil.Setloadini("c:\\gameinfo_temp.ini")

}

func main() {

	//getdict()

	//getdicttag()

	//getfileutil()

	//getiniutil()

	//getjsonutil()

	//getiniwinutil()

	//getregistryutil()

	//getstringutil()

	//gettimeutil()

	//getziputil()
}
