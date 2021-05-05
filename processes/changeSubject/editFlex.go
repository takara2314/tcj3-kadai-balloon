package changeSubject

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func editFlex(base []byte, class string) ([]byte, error) {
	var result string = string(base)
	var day int = int(time.Now().Weekday())
	if day == 0 || day == 6 {
		day = 5
	}

	switch class {
	case "J3A":
		fmt.Println(Config.Schedules.A[day-1])

		for period, subject := range Config.Schedules.A[day-1] {
			subject = strings.TrimPrefix(subject, "&")

			lecture := []flexJsonContent{
				flexJsonContent{
					Type:    "text",
					Text:    period,
					Margin:  "xl",
					Gravity: "center",
					Align:   "end",
					Size:    "sm",
				},
				flexJsonContent{
					Type: "button",
					Action: flexJsonContentAction{
						Type:  "message",
						Label: subject,
						Text:  subject + ",A",
					},
					Height: "sm",
					Flex:   3,
				},
			}

			lectureJson, err := json.Marshal(lecture)
			if err != nil {
				return nil, err
			}

			result = strings.Replace(
				result,
				`{"contents":"THIS"}`,
				string(lectureJson),
			)

		}

	case "J3B":
		fmt.Println(Config.Schedules.B[day-1])
	}
}
