package util

import (
	"strings"
	"time"
)

const (
	GMT_7 = 7 * 3600
)

var TimeZone *time.Location

func SetTimeZone(location string) {
	TimeZone, _ = time.LoadLocation(location)
}

func TimeNow() time.Time {
	return time.Now().In(TimeZone)
}

func ParseDateToUnix(date string) int {

	if strings.Contains(date, "-") {
		timeParse, err := time.Parse("2006-01-02 15:04", date)
		if err == nil {
			return int(timeParse.Unix()) - GMT_7
		}
	}

	if !strings.Contains(date, ":") {
		timeParse, err := time.Parse("02/01/2006", date)
		if err == nil {
			return int(timeParse.Unix()) - GMT_7
		}
	}

	timeParse, err := time.Parse("2/1/2006 15:04", date)
	if err == nil {
		return int(timeParse.Unix()) - GMT_7
	}

	timeParse, err = time.Parse("2006-01-02 15:04:05", date)
	if err == nil {
		return int(timeParse.Unix()) - GMT_7
	}

	return 0
}
