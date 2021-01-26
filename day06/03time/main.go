package main

import (
	"fmt"
	"time"
)

//时间
func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Date())

	//时间戳
	fmt.Println(now.Unix())

	//time.Unix()
	ret := time.Unix(1132566389, 0)
	fmt.Println(ret)
	fmt.Println(ret.Year())
	fmt.Println(ret.Month())
	fmt.Println(ret.Day())

	//时间间隔
	fmt.Println(time.Second)

	//now + 24小时
	fmt.Println(now.Add(24 * time.Hour))

	// //定时器
	// timer:=.tick
	// timer := time.Tick(time.Second) //定义一个1秒间隔的定时器
	// for i := range timer {
	// 	fmt.Println(i) //每秒都会执行的任务
	// }

	//时间格式化
	//2019-08-03
	fmt.Println(now.Format("2006-01-02"))
	//2021/01/26 06:10:44 PM
	fmt.Println(now.Format("2006/01/02 03:04:05 PM"))
	//2021/01/26 06:10:44.2345 PM	精确到毫秒
	fmt.Println(now.Format("2006/01/02 03:04:05.0000 PM"))

	//按照对应的格式解析字符串类型的时间
	timeObj, err := time.Parse("2006-01-02", "2000-08-09")
	if err != nil {
		fmt.Println("parse time failed,err%v\n", err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())
}
