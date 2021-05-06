package main

import (
	"log"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"

	"tcj3-kadai-tuika-kun/processes/addInfo"
	"tcj3-kadai-tuika-kun/processes/changeSubject"
	"tcj3-kadai-tuika-kun/processes/database"
)

func replyToStudent(event *linebot.Event, message string) {
	if strings.HasPrefix(message, "add") {
		err := addInfo.Response(bot, event, flexAddInfo)
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else if strings.HasPrefix(message, "change") {
		var err error
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
		var err error
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
		var err error
		var replyMessage string = "© 2021 Takara Hamaguchi"

		_, err = bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewTextMessage(replyMessage),
		).Do()
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else if message == "[提出期限を変更]" {
		var err error
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
		err := changeSubject.RegularResponse(
			bot,
			event,
			flexChangeSubject,
			"J3A",
		)
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else if isSubjectPrefix(message) {
		var err error
		var splited []string = strings.Split(message, ",")

		if len(splited) == 2 {
			var replyMessages []string = []string{
				"教科を設定しました。",

				"課題の提出期限を設定してください。",

				"" +
					"「XX月XX日 XX時XX分」と打てばその日時に、\n" +
					"「XX月XX日 XX限目」と打てばその時間が始まる日時に期限を設定できます。\n" +
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

		} else {
			var replyMessage string = "「教科,組」という形で送信してください。"

			_, err = bot.ReplyMessage(
				event.ReplyToken,
				linebot.NewTextMessage(replyMessage),
			).Do()
			if err != nil {
				log.Println(err)
				panic(err)
			}
		}

	} else if strings.HasPrefix(message, "others") {
		var err error
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
		err := changeSubject.RegularResponse(
			bot,
			event,
			flexChangeSubject,
			"A",
		)
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else if message == "B" {
		err := changeSubject.RegularResponse(
			bot,
			event,
			flexChangeSubject,
			"B",
		)
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else if strings.HasPrefix(message, "goodbye") {
		var err error
		var replyMessages []string = []string{
			"あなたのユーザーデータを消しました。",

			"また学籍番号を打てば再登録できます。",
		}

		err = database.RemoveUser(&dbCtx, dbClient, &users, event.Source.UserID)
		if err != nil {
			log.Println(err)
			panic(err)
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

	} else {
		var err error
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

func isSubjectPrefix(message string) bool {
	for _, subject := range config.Subjects {
		if strings.HasPrefix(message, subject) {
			return true
		}
	}

	return false
}
