package main

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

func replyTextMessage(event *linebot.Event, message string) {
	if _, exist := users[event.Source.UserID]; exist {
		replyToStudent(event, message)
	} else {
		replyToGuest(event, message)
	}
}
