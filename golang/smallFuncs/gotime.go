package main

import "time"

// get the date of last monday
func LastMonday() time.Time {
	todayUnixTime := time.Now().Unix()
	todayWeekday := time.Now().Weekday().String()
	var weekdayNum int

	switch todayWeekday {
	case "Monday":
		weekdayNum = 1
	case "TuesDay":
		weekdayNum = 2
	case "Wednesday":
		weekdayNum = 3
	case "ThursDay":
		weekdayNum = 4
	case "Friday":
		weekdayNum = 5
	case "Saturday":
		weekdayNum = 6
	case "Sunday":
		weekdayNum = 7
	default:
	}

	lastMonday := time.Unix(todayUnixTime-int64((6+weekdayNum)*86400), 0)
	return lastMonday
}
