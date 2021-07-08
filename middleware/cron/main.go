package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

var (
	ch chan int
)

func initChan() {
	ch = make(chan int)
}

func sendData() {
	//ch<-1
	println(1)
}

func main() {
	cron := cron.New()

	//initChan()
	// 从右到左依次是：星期几-月-日-时-分-秒 每天x时x分x秒 执行
	//cron.AddFunc("* 1/2 * * * *", sendData)

	cron.AddFunc("*/5 * * * * *", func() {//每隔5秒执行一次
		fmt.Println("a: ", time.Now().Format("2006-01-02 15:04:05"))
	})

	cron.AddFunc("0 */1 * * * ?", func() {//每隔1分钟执行一次
		fmt.Println("b: ", time.Now().Format("2006-01-02 15:04:05"))
	})

	//go func() {println(<-ch)}()

	//cron.AddFunc("30 58 10 * * *", func() {
	//	fmt.Println("并发2", time.Now())
	//})

	// cron.AddFunc("@every 1s", func() {
	// 	fmt.Println("Every second")
	// })

	//cron.AddFunc("@every 0h0m1s", func() {
	//	fmt.Println("Every 1s after 0 hour, 0 minutes, 1 seconds")
	//})

	// cron.AddFunc("* 10 * * * 4", func() {
	// 	fmt.Println("Every Thursday on ten Minutes")
	// })

	// cron.AddFunc("0 30 * * * *", func() {
	// 	fmt.Println("Every hour on the half hour")
	// })

	cron.Start()

	defer cron.Stop() // Stop the scheduler (does not stop any jobs already running).

	select {}
}

// Second: (uint64) 1,
// Minute: (uint64) 1073741824,
// Hour: (uint64) 9223372036871553023,
// Dom: (uint64) 9223372041149743102,
// Month: (uint64) 9223372036854783998,
// Dow: (uint64) 9223372036854775935

// Field name   | Mandatory? | Allowed values  | Allowed special characters
// ----------   | ---------- | --------------  | --------------------------
// Seconds      | Yes        | 0-59            | * / , -
// Minutes      | Yes        | 0-59            | * / , -
// Hours        | Yes        | 0-23            | * / , -
// Day of month | Yes        | 1-31            | * / , - ?
// Month        | Yes        | 1-12 or JAN-DEC | * / , -
// Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?

//   0      1       2          3          4        5        6
//   Sunday Monday Tuesday  Wednesday  Thursday  Friday  Saturday

// Entry                  | Description                                | Equivalent To
// -----                  | -----------                                | -------------
// @yearly (or @annually) | Run once a year, midnight, Jan. 1st        | 0 0 0 1 1 *
// @monthly               | Run once a month, midnight, first of month | 0 0 0 1 * *
// @weekly                | Run once a week, midnight between Sat/Sun  | 0 0 0 * * 0
// @daily (or @midnight)  | Run once a day, midnight                   | 0 0 0 * * *
// @hourly                | Run once an hour, beginning of hour        | 0 0 * * * *

// @every <duration>
// For example, "@every 1h30m10s" would indicate a schedule that activates after 1 hour, 30 minutes, 10 seconds,
// and then every interval after that.

//每一个域可出现的字符如下：
//Seconds:          可出现     ", - * /"     四个字符，有效范围为0-59的整数
//Minutes:          可出现     ", - * /"     四个字符，有效范围为0-59的整数
//Hours:            可出现     ", - * /"     四个字符，有效范围为0-23的整数
//DayofMonth:       可出现     ", - * / ? L W C"     八个字符，有效范围为0-31的整数
//Month:            可出现     ", - * /"     四个字符，有效范围为1-12的整数或JAN-DEc
//DayofWeek:        可出现     ", - * / ? L C #"     四个字符，有效范围为1-7的整数或SUN-SAT两个范围。1表示星期天，2表示星期一， 依次类推
//Year:             可出现     ", - * /"     四个字符，有效范围为1970-2099年


//https://www.cnblogs.com/zuxingyu/p/6023919.html

//cron特定字符说明
//　　1）星号(*)
//　　　　表示 cron 表达式能匹配该字段的所有值。如在第5个字段使用星号(month)，表示每个月
//
//　　2）斜线(/)
//　　　　表示增长间隔，如第1个字段(minutes) 值是 3-59/15，表示每小时的第3分钟开始执行一次，之后每隔 15 分钟执行一次（即 3、18、33、48 这些时间点执行），这里也可以表示为：3/15
//
//　　3）逗号(,)
//　　　　用于枚举值，如第6个字段值是 MON,WED,FRI，表示 星期一、三、五 执行
//
//　　4）连字号(-)
//　　　　表示一个范围，如第3个字段的值为 9-17 表示 9am 到 5pm 直接每个小时（包括9和17）
//
//　　5）问号(?)
//　　　　只用于 日(Day of month) 和 星期(Day of week)，表示不指定值，可以用于代替 *
//
//　　6）L，W，#
//　　　　Go中没有L，W，#的用法，下文作解释。
//

//cron举例说明
//每隔5秒执行一次：*/5 * * * * ?
//每隔1分钟执行一次：0 */1 * * * ?
//每天23点执行一次：0 0 23 * * ?
//每天凌晨1点执行一次：0 0 1 * * ?
//每月1号凌晨1点执行一次：0 0 1 1 * ?
//在26分、29分、33分执行一次：0 26,29,33 * * * ?
//每天的0点、13点、18点、21点都执行一次：0 0 0,13,18,21 * * ?

//每月最后一天23点执行一次：0 0 23 L * ?
//每周星期天凌晨1点实行一次：0 0 1 ? * L


//# ┌───────────── min (0 - 59)
//# │ ┌────────────── hour (0 - 23)
//# │ │ ┌─────────────── day of month (1 - 31)
//# │ │ │ ┌──────────────── month (1 - 12)
//# │ │ │ │ ┌───────────────── day of week (0 - 6) (0 to 6 are Sunday to
//# │ │ │ │ │                  Saturday, or use names; 7 is also Sunday)
//# │ │ │ │ │
//# │ │ │ │ │
//# * * * * *  command to execute