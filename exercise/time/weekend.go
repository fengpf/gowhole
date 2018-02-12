package main

import (
	"time"
)

// func BeginningOfHour() time.Time {
// 	return time.Time.Truncate(time.Hour)
// }

func beginningOfDay(t time.Time) time.Time {
	d := time.Duration(-t.Hour()) * time.Hour
	return t.Truncate(time.Hour).Add(d)
}

func getTuesday(now time.Time) time.Time {
	t := beginningOfDay(now)
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	d := time.Duration(-weekday+2) * 24 * time.Hour
	return t.Truncate(time.Hour).Add(d)
}

func getSunday(now time.Time) time.Time {
	t := beginningOfDay(now)
	weekday := int(t.Weekday())
	if weekday == 0 {
		return t
	}
	d := time.Duration(7-weekday) * 24 * time.Hour
	return t.Truncate(time.Hour).Add(d)
}

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	// var sd string
	// t := time.Now().AddDate(0, 0, -3).Add(-3 * time.Minute)
	// println(t.Format("2006-01-02 15:04:05"))
	// tdStr := getTuesday(t).Add(12 * time.Hour).Format("2006-01-02 15:04:05")
	// println(tdStr)
	// td := getTuesday(t).Add(12 * time.Hour)
	// if t.Before(td) {
	// 	sd = getSunday(t.AddDate(0, 0, -14)).Format("2006-01-02 15:04:05")
	// } else {
	// 	sd = getSunday(t.AddDate(0, 0, -7)).Format("2006-01-02 15:04:05")
	// }
	// println(sd)

	// mtStr := now.New(t).Monday().Add(12 * time.Hour).Format("2006-01-02 15:04:05")
	// mt := now.New(t).Monday().AddDate(0, 0, 1).Add(-12 * time.Hour)
	// if t.Before(mt) {
	// 	dt = now.New(t).EndOfSunday().AddDate(0, 0, -14).Format("2006-01-02 15:04:05")
	// } else {
	// 	dt = now.New(t).EndOfSunday().AddDate(0, 0, -7).Add(-12 * time.Hour).Format("2006-01-02 15:04:05")
	// }
}
