package main

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

func replyTextMessage(event *linebot.Event, message string) {
	replyToStudent(event, message)
}
