package setDue

import "time"

func thisWeekday(target time.Time, weekday int) time.Time {
	var duration int = weekday - int(target.Weekday())

	return target.Add(time.Duration(duration*24) * time.Hour)
}
