package setDue

import (
	"strconv"
	"time"

	"golang.org/x/text/unicode/norm"
)

func SetDue(data []string) (due time.Time) {
	var today time.Time = time.Now()

	if data[9] == "" {
		data[9] = strconv.Itoa(int(today.Weekday()))
	}

	if data[2] != "" {
		month, _ := strconv.Atoi(
			string(norm.NFKC.Bytes([]byte(data[2]))),
		)
		day, _ := strconv.Atoi(
			string(norm.NFKC.Bytes([]byte(data[5]))),
		)

		due = time.Date(
			today.Year(),
			time.Month(month),
			day,
			today.Hour(),
			today.Minute(),
			0,
			0,
			loc,
		)

		if today.After(due) {
			due = time.Date(
				due.Year()+1,
				due.Month(),
				due.Day(),
				due.Hour(),
				due.Minute(),
				0,
				0,
				loc,
			)
		}

	} else if data[1] == "今日" || data[1] == "本日" {
		due = today

	} else if data[1] == "明日" {
		due = today.Add(time.Duration(1*24) * time.Hour)

	} else if data[1] == "明後日" {
		due = today.Add(time.Duration(2*24) * time.Hour)

	} else if data[1] == "明々後日" {
		due = today.Add(time.Duration(3*24) * time.Hour)

	} else {
		weekdays := map[string]int{
			"日": 0,
			"月": 1,
			"火": 2,
			"水": 3,
			"木": 4,
			"金": 5,
			"土": 6,
		}

		switch data[8] {
		case "今週":
			due = thisWeekday(today, weekdays[data[9]])
		case "来週":
			due = nextWeekday(today, weekdays[data[9]])
		case "再来週":
			due = nextWeekday(
				nextWeekday(today, weekdays[data[9]]),
				weekdays[data[9]],
			)
		default:
			due = thisWeekdayOnlyFuture(today, weekdays[data[9]])
		}
	}

	if data[13] != "" {
		hour, _ := strconv.Atoi(
			string(norm.NFKC.Bytes([]byte(data[13]))),
		)
		minute, _ := strconv.Atoi(
			string(norm.NFKC.Bytes([]byte(data[16]))),
		)

		due = time.Date(
			due.Year(),
			due.Month(),
			due.Day(),
			hour,
			minute,
			0,
			0,
			loc,
		)

	} else {
		period := string(norm.NFKC.Bytes([]byte(data[10]))) + "限目"
		classStart, _ := time.Parse(
			"15:04",
			Config.Timetables[period][0].(string),
		)

		due = time.Date(
			due.Year(),
			due.Month(),
			due.Day(),
			classStart.Hour(),
			classStart.Minute(),
			0,
			0,
			loc,
		)
	}

	return
}
