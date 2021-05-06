package addInfo

import (
	"strings"
	"time"
)

func nearTimetable(class string, weekday int) (subject string, targetClass string, classStart time.Time) {
	var today time.Time = time.Now()

	var schedules [][][]string
	switch class {
	case "J3A":
		schedules = Config.Schedules.A
	case "J3B":
		schedules = Config.Schedules.B
	}

	var matched bool = false

	for _, configSubject := range schedules[weekday-1] {
		var classLong int = Config.Timetables[configSubject[0]][1].(int)

		start, _ := time.Parse(
			"15:04",
			Config.Timetables[configSubject[0]][0].(string),
		)

		classStart = time.Date(
			today.Year(),
			today.Month(),
			today.Day(),
			start.Hour(),
			start.Minute(),
			0,
			0,
			loc,
		)

		var classBeforeEnd time.Time = classStart.
			Add(time.Duration(classLong-30) * time.Minute)

		if today.After(classBeforeEnd) {
			matched = true

			if strings.HasPrefix(configSubject[1], "&") {
				subject = strings.TrimPrefix(configSubject[1], "&")
				targetClass = "J3A, J3B"
			} else {
				subject = configSubject[1]
				targetClass = class
			}
		}
	}

	if !matched {
		var oldday time.Time = time.Now()

		switch weekday {
		case 0:
			weekday = 6
			oldday = oldday.Add(time.Duration(-2*24) * time.Hour)

		case 1:
			weekday = 6
			oldday = oldday.Add(time.Duration(-3*24) * time.Hour)

		default:
			oldday = oldday.Add(time.Duration(-1*24) * time.Hour)
		}

		lecture_num := len(schedules[weekday-1])

		start, _ := time.Parse(
			"15:04",
			Config.Timetables[schedules[weekday-1][lecture_num-1][0]][0].(string),
		)

		classStart = time.Date(
			oldday.Year(),
			oldday.Month(),
			oldday.Day(),
			start.Hour(),
			start.Minute(),
			0,
			0,
			loc,
		)

		if strings.HasPrefix(schedules[weekday-1][lecture_num-1][1], "&") {
			subject = strings.TrimPrefix(schedules[weekday-1][lecture_num-1][1], "&")
			targetClass = "J3A, J3B"
		} else {
			subject = schedules[weekday-1][lecture_num-1][1]
			targetClass = class
		}
	}

	return
}
