package addInfo

import (
	"strconv"
	"strings"
	"time"
)

func editFlex(base []byte, class string) ([]byte, error) {
	var result string = string(base)

	var weekday int = int(time.Now().Weekday())

	subject, targetClass, classStart := nearTimetable(class, weekday)
	var due time.Time = surmiseDue(subject, classStart)

	result = strings.Replace(
		result,
		"${subject}",
		subject,
		1,
	)
	result = strings.Replace(
		result,
		"${class}",
		targetClass,
		1,
	)
	result = strings.Replace(
		result,
		"${month}",
		strconv.Itoa(int(due.Month())),
		1,
	)
	result = strings.Replace(
		result,
		"${day}",
		strconv.Itoa(due.Day()),
		1,
	)
	result = strings.Replace(
		result,
		"${weekday}",
		strconv.Itoa(int(due.Weekday())),
		1,
	)
	result = strings.Replace(
		result,
		"${hour}",
		strconv.Itoa(due.Hour()),
		1,
	)
	result = strings.Replace(
		result,
		"${minute}",
		strconv.Itoa(due.Minute()),
		1,
	)

	return []byte(result), nil
}
