package setDue

import "time"

func thisWeekdayOnlyFuture(target time.Time, weekday int) time.Time {
	var duration int = weekday - int(target.Weekday())

	if duration < 0 {
		return target.Add(time.Duration((duration+7)*24) * time.Hour)
	}

	return target.Add(time.Duration(duration*24) * time.Hour)
}
