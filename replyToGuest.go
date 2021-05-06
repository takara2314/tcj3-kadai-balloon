package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"tcj3-kadai-tuika-kun/processes/database"

	"github.com/line/line-bot-sdk-go/linebot"
	"golang.org/x/text/unicode/norm"
)

func replyToGuest(event *linebot.Event, message string) {
	studentNumReg := regexp.MustCompile(`^([0-9]|[０-９]){5}$`)
	authCodeReg := regexp.MustCompile(`^KTK[0-9]{5}$`)
	aClassReg := regexp.MustCompile(`^(A|Ａ).{0,2}(組|ぐみ|グミ|ｸﾞﾐ|class|).{0,5}$`)
	bClassReg := regexp.MustCompile(`^(B|Ｂ).{0,2}(組|ぐみ|グミ|ｸﾞﾐ|class|).{0,5}$`)
	tmpReg := regexp.MustCompile(`^tmp (add|remove)( |)(A|B|)( |)([0-9]{5}|)$`)

	if studentNumReg.MatchString(message) {
		var err error
		var replyMessages []string = []string{
			fmt.Sprintf(
				"%s@toba-cmt.ac.jp にメールを送信しました。",
				string(norm.NFKC.Bytes([]byte(message))),
			),

			"メールに書かれている認証コードを僕に送信してください。",
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

	} else if authCodeReg.MatchString(message) {
		var err error
		var replyMessages []string = []string{
			"認証に成功しました！",

			"" +
				"あなたは何組ですか？\n" +
				"A組なら「A」、B組なら「B」と送信してください。",
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

	} else if aClassReg.MatchString(message) {
		var err error
		var replyMessages []string = []string{
			"A組ですね！わかりました。",

			"これにて僕を使えるようになりました。",

			"[僕の使い方]",
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

	} else if bClassReg.MatchString(message) {
		var err error
		var replyMessages []string = []string{
			"B組ですね！わかりました。",

			"これにて僕を使えるようになりました。",

			"[僕の使い方]",
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

	} else if tmpReg.MatchString(message) {
		var err error
		var replyMessage string
		var splited []string = strings.Split(message, " ")

		switch splited[1] {
		case "add":
			profile, err := bot.GetProfile(event.Source.UserID).Do()
			if err != nil {
				log.Println(err)
				panic(err)
			}

			replyMessage = "Added."

			studentNum, _ := strconv.Atoi(splited[3])

			var class string
			if splited[2] == "A" {
				class = "J3A"
			} else {
				class = "J3B"
			}

			err = database.AddUser(&dbCtx, dbClient, &users, &map[string]interface{}{
				"line_id":     event.Source.UserID,
				"line_name":   profile.DisplayName,
				"student_num": studentNum,
				"class":       class,
			})
			if err != nil {
				log.Println(err)
				panic(err)
			}

		case "remove":
			replyMessage = "Removed. (Not system synced.)"
		}

		_, err = bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewTextMessage(replyMessage),
		).Do()
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else {
		var err error
		var replyMessages []string = []string{
			"鳥羽商船の学生以外は使用できません。",

			"" +
				"「19299」と学籍番号のみ送信し、認証を開始しましょう。\n" +
				"この情報は認証以外では一切使用しませんので、ご安心ください。",
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
	}
}
