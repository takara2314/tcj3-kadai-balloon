package addInfo

import "github.com/line/line-bot-sdk-go/linebot"

func Response(bot *linebot.Client, event *linebot.Event, flexAddInfo []byte) error {
	var err error
	var replyMessage string
	var res *linebot.UserProfileResponse
	var flexMessage linebot.FlexContainer

	res, err = bot.GetProfile(event.Source.UserID).Do()
	if err != nil {
		return err
	}

	flexMessage, err = linebot.UnmarshalFlexMessageJSON(flexAddInfo)
	if err != nil {
		return err
	}

	replyMessage = "課題のタイトルを教えてください。 " + res.DisplayName + " (" + event.Source.UserID + ")"

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
