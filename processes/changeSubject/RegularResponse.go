package changeSubject

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/linebot"
)

func RegularResponse(bot *linebot.Client, event *linebot.Event, flexChangeSubject []byte, class string) error {
	var err error
	var flexMessage linebot.FlexContainer

	flex, err := editFlex(flexChangeSubject, class)
	if err != nil {
		return err
	}

	fmt.Println(flex)

	flexMessage, err = linebot.UnmarshalFlexMessageJSON(flex)
	if err != nil {
		fmt.Println("エラーが起こったよ！")
		return err
	}

	_, err = bot.ReplyMessage(
		event.ReplyToken,
		linebot.NewFlexMessage("教科を選んでください。", flexMessage),
	).Do()
	if err != nil {
		fmt.Println("ERROR LOVE!")
		return err
	}

	return nil
}
