// main
package main

import (
	"Daseonglib/stringutil"
	"Daseonglib/timeutil"
	"fmt"
)

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

func main() {

	getstringutil()

	gettimeutil()
}
