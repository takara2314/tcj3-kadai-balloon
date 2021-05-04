package main

import (
	"log"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

func replyTextMessage(event *linebot.Event, message string) {
	var err error
	var replyMessage string

	if strings.HasPrefix(message, "add") {
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
			linebot.NewFlexMessage("追加する課題情報", flexMessage),
			linebot.NewTextMessage(replyMessage),
		).Do()
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else if strings.HasPrefix(message, "change") {
		replyMessage = "change OK"

		_, err = bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewTextMessage(replyMessage),
		).Do()
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else if strings.HasPrefix(message, "remove") {
		replyMessage = "remove OK"

		_, err = bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewTextMessage(replyMessage),
		).Do()
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else if strings.HasPrefix(message, "info") {
		replyMessage = "info OK"

		_, err = bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewTextMessage(replyMessage),
		).Do()
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else {
		replyMessage = "？"

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
