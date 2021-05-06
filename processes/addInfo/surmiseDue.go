package addInfo

import "time"

func surmiseDue(subject string, classStart time.Time) time.Time {
	switch subject {
	case "キャリアデザイン１":
		classStart = classStart.Add(time.Duration(4*24) * time.Hour)
		classStart = time.Date(
			classStart.Year(),
			classStart.Month(),
			classStart.Day(),
			17,
			0,
			0,
			0,
			loc,
		)

	case "工学数理基礎１":
		classStart = time.Date(
			classStart.Year(),
			classStart.Month(),
			classStart.Day(),
			23,
			59,
			0,
			0,
			loc,
		)

	case "プログラミング２":
		classStart = classStart.Add(time.Duration(7*24) * time.Hour)
		classStart = time.Date(
			classStart.Year(),
			classStart.Month(),
			classStart.Day(),
			13,
			0,
			0,
			0,
			loc,
		)

	case "情報工学３":
		classStart = classStart.Add(time.Duration(7*24) * time.Hour)
		classStart = time.Date(
			classStart.Year(),
			classStart.Month(),
			classStart.Day(),
			23,
			59,
			0,
			0,
			loc,
		)

	default:
		classStart = classStart.Add(time.Duration(7*24) * time.Hour)
	}

	return classStart
}
