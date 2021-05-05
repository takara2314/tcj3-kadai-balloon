package main

import (
	"context"
	"io/ioutil"
	"log"
	"tcj3-kadai-tuika-kun/processes/changeSubject"
	"tcj3-kadai-tuika-kun/processes/database"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/line/line-bot-sdk-go/linebot"
	"gopkg.in/yaml.v2"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"

	"tcj3-kadai-tuika-kun/types"
)

var (
	bot               *linebot.Client
	dbCtx             context.Context
	dbClient          *firestore.Client
	config            types.ConfigYaml
	flexAddInfo       []byte
	flexChangeSubject []byte
	users             map[string][]interface{} = make(map[string][]interface{}, 128)
	kadais            map[string][]interface{} = make(map[string][]interface{}, 1024)
)

func init() {
	var err error

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		loc = time.FixedZone("Asia/Tokyo", 9*60*60)
	}
	time.Local = loc

	configData, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	changeSubject.Config = &config

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
