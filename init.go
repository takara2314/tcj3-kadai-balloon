package main

import (
	"io/ioutil"
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

var (
	bot               *linebot.Client
	subjectList       []string
	flexAddInfo       []byte
	flexChangeSubject []byte
)

func init() {
	var err error

	subjectList = []string{
		"国語３",
		"現代社会",
		"日本語教育１",
		"日本語教育２",
		"微分積分２",
		"代数・幾何２",
		"化学",
		"保健体育３",
		"Level Up English１",
		"Level Up English２",
		"プログラミング２",
		"情報工学３",
		"WEBアプリケーション",
		"電気電子回路",
		"工業力学１",
		"材料学",
		"機械製図",
		"機械加工実習",
		"マイコン工学",
		"計測工学",
		"工学数理基礎１",
		"工学数理基礎２",
		"キャリアデザイン１",
		"PBL３",
	}

	flexAddInfo, err = ioutil.ReadFile("./templates/addInfo.json")
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	flexChangeSubject, err = ioutil.ReadFile("./templates/changeSubject.json")
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
}
