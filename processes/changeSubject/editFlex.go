package changeSubject

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

func editFlex(base []byte, class string) ([]byte, error) {
	var result string = string(base)

	var classLetter string
	switch class {
	case "J3A":
		classLetter = "A"
	case "J3B":
		classLetter = "B"
	}

	var weekday int = int(time.Now().Weekday())
	if weekday == 0 || weekday == 6 {
		weekday = 5
	}

	result = strings.Replace(
		result,
		"${class}",
		class,
		1,
	)

	_, month, day := nearThisWeekday(weekday)
	result = strings.Replace(
		result,
		"${month}",
		strconv.Itoa(month),
		1,
	)
	result = strings.Replace(
		result,
		"${day}",
		strconv.Itoa(day),
		1,
	)
	result = strings.Replace(
		result,
		"${weekday}",
		[]string{"日", "月", "火", "水", "木", "金", "土"}[weekday],
		1,
	)

	var lectures []flexJsonContent
	var schedule *map[string]string
	switch class {
	case "J3A":
		schedule = &Config.Schedules.A[weekday-1]
	case "J3B":
		schedule = &Config.Schedules.B[weekday-1]
	}

	for period, subject := range *schedule {
		subject = strings.TrimPrefix(subject, "&")

		lecture := flexJsonContent{
			Type:   "box",
			Layout: "horizontal",
			Contents: []flexJsonContent2nd{
				{
					Type:    "text",
					Text:    period,
					Margin:  "xl",
					Gravity: "center",
					Align:   "end",
					Size:    "sm",
				},
				{
					Type: "button",
					Action: flexJsonContent2ndAction{
						Type:  "message",
						Label: subject,
						Text:  subject + "," + classLetter,
					},
					Height: "sm",
					Flex:   3,
				},
			},
		}

		lectures = append(lectures, lecture)
	}

	flexContent := flexJson{
		Contents: lectures,
	}

	flexContentBytes, err := json.Marshal(flexContent)
	if err != nil {
		return nil, err
	}

	flexContentStr := string(flexContentBytes)
	flexContentStr = strings.TrimPrefix(flexContentStr, "{")
	flexContentStr = strings.TrimSuffix(flexContentStr, "}")

	result = strings.Replace(
		result,
		`"contents": "${contents}"`,
		flexContentStr,
		1,
	)

	return []byte(result), nil
}
