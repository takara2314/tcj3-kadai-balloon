package main

import (
	"context"
	"io/ioutil"
	"log"
	"tcj3-kadai-tuika-kun/processes/database"

	"cloud.google.com/go/firestore"
	"github.com/line/line-bot-sdk-go/linebot"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var (
	bot         *linebot.Client
	dbCtx       context.Context
	dbClient    *firestore.Client
	subjectList []string = []string{
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
	flexAddInfo       []byte
	flexChangeSubject []byte
	users             map[string][]interface{} = make(map[string][]interface{}, 128)
	kadais            map[string][]interface{} = make(map[string][]interface{}, 1024)
)

func init() {
	var err error

	flexAddInfo, err = ioutil.ReadFile("./templates/addInfo.json")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	flexChangeSubject, err = ioutil.ReadFile("./templates/changeSubject.json")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	dbCtx = context.Background()
	sa := option.WithCredentialsFile("./tcj3-kadai-tuika-kun-firebase-adminsdk-mqxg5-96023eb6d4.json")
	app, err := firebase.NewApp(dbCtx, nil, sa)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	dbClient, err = app.Firestore(dbCtx)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer dbClient.Close()

	err = initInfo()
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func initInfo() error {
	var err error

	err = database.GetUsers(&dbCtx, dbClient, &users)
	if err != nil {
		return err
	}

	err = database.GetKadais(&dbCtx, dbClient, &kadais)
	if err != nil {
		return err
	}

	return nil
}
