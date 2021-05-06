package setDue

import (
	"time"
)

func nextWeekday(target time.Time, weekday int) time.Time {
	var duration int = weekday - int(target.Weekday())

	return target.Add(time.Duration((duration+7)*24) * time.Hour)
}
