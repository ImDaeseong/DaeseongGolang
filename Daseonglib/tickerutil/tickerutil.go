package tickerutil

import (
	"Daseonglib/winapiutil"
	"fmt"
	"time"
)

func StopWatch_second(startSecond, endSecond time.Duration) {

	ticker := time.NewTicker(startSecond * time.Second)

	go printSecond(ticker)

	time.Sleep(endSecond * time.Second)
	ticker.Stop()
}

func printSecond(ticker *time.Ticker) {

	index := 1
	for _ = range ticker.C {
		fmt.Println(index)
		index++
	}
}

func StopWatch_time(startSecond, endSecond time.Duration) {

	startticker := time.NewTicker(startSecond * time.Second)
	stoptime := time.After(endSecond * time.Second)
	defer startticker.Stop()

	for {
		select {
		case <-startticker.C:
			sResult := fmt.Sprintf("%02d:%02d:%02d", time.Now().Hour(), time.Now().Minute(), time.Now().Second())
			fmt.Println(sResult)
		case <-stoptime:
			//리턴하면 종료됨
			return
		}
	}
}

func StopWatch_multi(startSecond, endSecond time.Duration) {

	Ticker1 := time.NewTicker(startSecond * time.Second)
	Ticker2 := time.NewTicker(startSecond * time.Second)
	Ticker3 := time.NewTicker(startSecond * time.Second)
	Ticker4 := time.NewTicker(startSecond * time.Second)
	defer Ticker1.Stop()
	defer Ticker2.Stop()
	defer Ticker3.Stop()
	defer Ticker4.Stop()

	bDoneChan := make(chan bool)

	//종료
	go func() {
		time.Sleep(endSecond * time.Second)

		//select case bDoneChan = true 채널에 입력
		bDoneChan <- true
	}()

	for {
		select {
		case <-Ticker1.C:

			sResult := fmt.Sprintf("Ticker1 - %02d:%02d:%02d", time.Now().Hour(), time.Now().Minute(), time.Now().Second())
			fmt.Println(sResult)

		case <-Ticker2.C:

			sResult := fmt.Sprintf("Ticker2 - %02d:%02d:%02d", time.Now().Hour(), time.Now().Minute(), time.Now().Second())
			fmt.Println(sResult)

		case <-Ticker3.C:

			sResult := fmt.Sprintf("Ticker3 - %02d:%02d:%02d", time.Now().Hour(), time.Now().Minute(), time.Now().Second())
			fmt.Println(sResult)

		case <-Ticker4.C:

			sResult := fmt.Sprintf("Ticker4 - %02d:%02d:%02d", time.Now().Hour(), time.Now().Minute(), time.Now().Second())
			fmt.Println(sResult)

		case <-bDoneChan:
			//리턴하면 종료됨
			return
		}
	}

}

func Callurl(URLS []string, timeSecond time.Duration) {

	/*
		for v := range URLS {
			fmt.Println("url:", URLS[v])
		}
	*/

	Ticker := time.NewTicker(timeSecond * time.Second)
	defer Ticker.Stop()

	nIndex := 0
	for {
		select {
		case <-Ticker.C:

			if nIndex > (len(URLS) - 1) {
				nIndex = 0
			}
			//fmt.Println("url:", nIndex, URLS[nIndex])

			winapiutil.ShellExecute(0, "open", URLS[nIndex], "", "", winapiutil.SW_SHOW)

			nIndex++
		}
	}

}
