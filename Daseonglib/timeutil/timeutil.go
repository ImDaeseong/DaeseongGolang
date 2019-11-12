// timeutil
package timeutil

import (
	"errors"
	"fmt"
	"time"
)

func GetFullCurrentDay() string {

	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	hour := time.Now().Hour()
	minute := time.Now().Minute()
	second := time.Now().Second()

	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)
}

func GetCurrentDay() string {

	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()

	return fmt.Sprintf("%04d-%02d-%02d", year, month, day)
}

func GetCurrentTime() string {

	hour := time.Now().Hour()
	minute := time.Now().Minute()
	second := time.Now().Second()

	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
}

func GetToday() string {

	sResult := ""

	weekday := time.Now().Weekday()

	//fmt.Println(weekday)
	//fmt.Println(int(weekday))

	switch weekday {
	case time.Sunday:
		sResult = fmt.Sprintf("일요일")
	case time.Monday:
		sResult = fmt.Sprintf("월요일")
	case time.Tuesday:
		sResult = fmt.Sprintf("화요일")
	case time.Wednesday:
		sResult = fmt.Sprintf("수요일")
	case time.Thursday:
		sResult = fmt.Sprintf("목요일")
	case time.Friday:
		sResult = fmt.Sprintf("금요일")
	case time.Saturday:
		sResult = fmt.Sprintf("토요일")
	}

	return sResult
}

func GetYear() string {

	now := time.Now()

	return fmt.Sprintf("%04d", now.Year())
}

func GetMonth() string {

	now := time.Now()

	return fmt.Sprintf("%02d", now.Month())
}

func GetDay() string {

	now := time.Now()

	return fmt.Sprintf("%02d", now.Day())
}

func SetDay(nDay int) string {

	now := time.Now()

	var day time.Time

	day = now.AddDate(0, 0, nDay)
	sDay := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", day.Year(), day.Month(), day.Day(), day.Hour(), day.Minute(), day.Second())
	return sDay
}

func SetMonth(nMonth int) string {

	now := time.Now()

	var month time.Time

	month = now.AddDate(0, nMonth, 0)
	sMonth := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", month.Year(), month.Month(), month.Day(), month.Hour(), month.Minute(), month.Second())
	return sMonth
}

func SetYear(nYear int) string {

	now := time.Now()

	var year time.Time

	year = now.AddDate(nYear, 0, 0)
	sYear := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", year.Year(), year.Month(), year.Day(), year.Hour(), year.Minute(), year.Second())
	return sYear
}

//string to time
func ConvertStrToTime(sTime string) (time.Time, error) {

	tm, err := time.Parse("2006-01-02 15:04:05", sTime)
	if err == nil {
		return tm, nil
	}

	return time.Now(), errors.New("")
}

func SubstringTime(sStart, sEnd string) float64 {

	starttm, _ := time.Parse("2006-01-02 15:04:05", sStart)
	endtm, _ := time.Parse("2006-01-02 15:04:05", sEnd)
	diff := endtm.Sub(starttm)

	//hour := int(diff.Hours())
	//minutes := int(diff.Minutes())
	//seconds := int(diff.Seconds())
	//days := int(diff.Hours() / 24)
	//fmt.Println("시:" + fmt.Sprintf("%d", hour))
	//fmt.Println("분:" + fmt.Sprintf("%d", minutes))
	//fmt.Println("초:" + fmt.Sprintf("%d", seconds))
	//fmt.Println("일:" + fmt.Sprintf("%d", days))

	//minutes := int(diff.Seconds() / 60)
	//hour := int(minutes / 60)
	//fmt.Println("초를 분으로 :" + fmt.Sprintf("%d", minutes))
	//fmt.Println("분를 시으로 :" + fmt.Sprintf("%d", hour))

	return diff.Seconds()
}

func SubtimeTime(sStart, sEnd time.Time) float64 {

	diff := sEnd.Sub(sStart)

	return diff.Seconds()
}
