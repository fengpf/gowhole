package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

func main() {
	cron := cron.New()

	// cron.AddFunc("@every 1s", func() {
	// 	fmt.Println("Every second")
	// })

	cron.AddFunc("@every 0h0m1s", func() {
		fmt.Println("Every 1s after 0 hour, 0 minutes, 1 seconds")
	})

	// cron.AddFunc("* 10 * * * 4", func() {
	// 	fmt.Println("Every Thursday on ten Minutes")
	// })

	// cron.AddFunc("0 30 * * * *", func() {
	// 	fmt.Println("Every hour on the half hour")
	// })

	cron.Start()

	time.Sleep(1000 * time.Second)
	cron.Stop() // Stop the scheduler (does not stop any jobs already running).
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
