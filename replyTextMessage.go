package main

import (
	"log"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

func replyTextMessage(event *linebot.Event, message string) {
	var err error

	if strings.HasPrefix(message, "add") {
		var replyMessage string
		var res *linebot.UserProfileResponse
		var flexMessage linebot.FlexContainer

		res, err = bot.GetProfile(event.Source.UserID).Do()
		if err != nil {
			log.Println(err)
			panic(err)
		}

		flexMessage, err = linebot.UnmarshalFlexMessageJSON(flexAddInfo)
		if err != nil {
			log.Println(err)
			panic(err)
		}

		replyMessage = "課題のタイトルを教えてください。 " + res.DisplayName + " (" + event.Source.UserID + ")"

		_, err = bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewFlexMessage("課題のタイトルを教えてください。", flexMessage),
			linebot.NewTextMessage(replyMessage),
		).Do()
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else if strings.HasPrefix(message, "change") {
		var replyMessage string = "change OK"

		_, err = bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewTextMessage(replyMessage),
		).Do()
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else if strings.HasPrefix(message, "remove") {
		var replyMessage string = "remove OK"

		_, err = bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewTextMessage(replyMessage),
		).Do()
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else if strings.HasPrefix(message, "info") {
		var replyMessage string = "info OK"

		_, err = bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewTextMessage(replyMessage),
		).Do()
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else if message == "[提出期限を変更]" {
		var replyMessages []string = []string{
			"課題の提出期限を設定してください。",

			"" +
				"「XX月XX日 XX時XX分」と打てばその日時に、「XX月XX日 XX限目」と打てばその時間が始まる日時に期限を設定できます。\n" +
				"「XX月XX日」の代わりに「来週」や「今週○曜日」にすることもできます。",
		}

		_, err = bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewTextMessage(replyMessages[0]),
			linebot.NewTextMessage(replyMessages[1]),
		).Do()
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else if message == "[教科を変更]" {
		var flexMessage linebot.FlexContainer

		flexMessage, err = linebot.UnmarshalFlexMessageJSON(flexChangeSubject)
		if err != nil {
			log.Println(err)
			panic(err)
		}

		_, err = bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewFlexMessage("教科を選んでください。", flexMessage),
		).Do()
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else if isSubject(message) {
		var replyMessages []string = []string{
			"教科を設定しました。",

			"課題の提出期限を設定してください。",

			"" +
				"「XX月XX日 XX時XX分」と打てばその日時に、「XX月XX日 XX限目」と打てばその時間が始まる日時に期限を設定できます。\n" +
				"「XX月XX日」の代わりに「来週」や「今週○曜日」にすることもできます。",
		}

		_, err = bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewTextMessage(replyMessages[0]),
			linebot.NewTextMessage(replyMessages[1]),
			linebot.NewTextMessage(replyMessages[2]),
		).Do()
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else if strings.HasPrefix(message, "others") {
		var splited []string = strings.Split(message, ",")

		if len(splited) == 2 {
			var flexMessage linebot.FlexContainer

			flexMessage, err = linebot.UnmarshalFlexMessageJSON(flexChangeSubject)
			if err != nil {
				log.Println(err)
				panic(err)
			}

			_, err = bot.ReplyMessage(
				event.ReplyToken,
				linebot.NewFlexMessage("教科を選んでください。", flexMessage),
			).Do()
			if err != nil {
				log.Println(err)
				panic(err)
			}

		} else {
			var replyMessage string = "「others,組」という形で送信してください。"

			_, err = bot.ReplyMessage(
				event.ReplyToken,
				linebot.NewTextMessage(replyMessage),
			).Do()
			if err != nil {
				log.Println(err)
				panic(err)
			}
		}

	} else if message == "A" {
		var flexMessage linebot.FlexContainer

		flexMessage, err = linebot.UnmarshalFlexMessageJSON(flexChangeSubject)
		if err != nil {
			log.Println(err)
			panic(err)
		}

		_, err = bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewFlexMessage("教科を選んでください。", flexMessage),
		).Do()
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else if message == "B" {
		var flexMessage linebot.FlexContainer

		flexMessage, err = linebot.UnmarshalFlexMessageJSON(flexChangeSubject)
		if err != nil {
			log.Println(err)
			panic(err)
		}

		_, err = bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewFlexMessage("教科を選んでください。", flexMessage),
		).Do()
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else {
		var replyMessage string = "？"

		_, err = bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewTextMessage(replyMessage),
		).Do()
		if err != nil {
			log.Println(err)
			panic(err)
		}
	}
}

func isSubject(subject string) bool {
	for _, str := range subjectList {
		if str == subject {
			return true
		}
	}

	return false
}
