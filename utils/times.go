package utils

import "time"

func GetCurrentTimeTimeZone(timeZone string) string {
	theTime := time.Now()
	loc, _ := time.LoadLocation(timeZone)
	theTime = theTime.In(loc)
	time := theTime.Format("[02 January 2006] 15:04:05 MST")
	return time
}
