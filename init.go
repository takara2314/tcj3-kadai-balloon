package main

import (
	"io/ioutil"
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

var (
	bot         *linebot.Client
	flexAddInfo []byte
)

func init() {
	var err error

	flexAddInfo, err = ioutil.ReadFile("./templates/addInfo.json")
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
}
