package utils

import (
	"fmt"
	"time"
)

// Returns the date of when with hours and seconds removed (set to 0).
func GetDay(when time.Time) time.Time {
	return time.Date(when.Year(), when.Month(), when.Day(), 0, 0, 0, 0, when.Location())
}

// Returns a timestamp in the form: YYYY-MM-DD
// If addSeconds is true, it adds HH:MM:SS
func GetUnixTimeStr(when time.Time, addSeconds bool) string {
	res := fmt.Sprintf("%02d-%02d-%02d", when.Year(), when.Month(), when.Day())
	if addSeconds {
		res += fmt.Sprintf(" %02d:%02d:%02d", when.Hour(), when.Minute(), when.Second())
	}
	return res
}
