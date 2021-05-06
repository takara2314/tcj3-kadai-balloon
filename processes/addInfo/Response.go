package addInfo

import "github.com/line/line-bot-sdk-go/linebot"

func Response(bot *linebot.Client, event *linebot.Event, flexAddInfo []byte, class string) error {
	var err error
	var replyMessage string
	var flexMessage linebot.FlexContainer

	flex, err := editFlex(flexAddInfo, class)
	if err != nil {
		return err
	}

	flexMessage, err = linebot.UnmarshalFlexMessageJSON(flex)
	if err != nil {
		return err
	}

	replyMessage = "課題のタイトルを教えてください。 "

	_, err = bot.ReplyMessage(
		event.ReplyToken,
		linebot.NewFlexMessage("課題のタイトルを教えてください。", flexMessage),
		linebot.NewTextMessage(replyMessage),
	).Do()
	if err != nil {
		return err
	}

	return nil
}
