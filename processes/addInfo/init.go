package addInfo

import (
	"tcj3-kadai-tuika-kun/types"
	"time"
)

var (
	Config *types.ConfigYaml
	loc    *time.Location
)

func init() {
	var err error
	loc, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		loc = time.FixedZone("Asia/Tokyo", 9*60*60)
	}
}
