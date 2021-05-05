package changeSubject

import "time"

func nearThisWeekday(weekday int) (year int, month int, day int) {
	var today time.Time = time.Now()
	var correctWeekday int = int(today.Weekday())

	if correctWeekday == weekday {
		year = today.Year()
		month = int(today.Month())
		day = today.Day()
		return
	}

	var near time.Time = time.Now()

	var durationDay int = 0
	if weekday == 0 {
		durationDay = 2
	} else if weekday == 6 {
		durationDay = 1
	}

	near = near.Add(time.Duration(-1*durationDay) * 24 * time.Hour)

	year = near.Year()
	month = int(near.Month())
	day = near.Day()
	return
}
