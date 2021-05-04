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
		res, err = bot.GetProfile(event.Source.UserID).Do()
		if err != nil {
			log.Println(err)
			panic(err)
		}

		replyMessage = "add OK " + res.DisplayName + " (" + event.Source.UserID + ")"

	} else if strings.HasPrefix(message, "change") {
		replyMessage = "change OK"

	} else if strings.HasPrefix(message, "remove") {
		replyMessage = "remove OK"

	} else if strings.HasPrefix(message, "info") {
		replyMessage = "info OK"

	} else {
		replyMessage = "？"
	}

	_, err = bot.ReplyMessage(
		event.ReplyToken,
		linebot.NewTextMessage(replyMessage),
	).Do()
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
