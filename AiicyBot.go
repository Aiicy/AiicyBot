package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func Separator() string {
	var path string = ""
	if os.IsPathSeparator('\\') {
		path = "\\"
	} else {
		path = "/"
	}

	return path
}

func GenRandomPicFromPath(path string) string {
	f, err := os.Stat(path)
	if err != nil {
		log.Fatalln(err)
	}
	if f.IsDir() != true {
		log.Fatalln(path, "is not a folder")
	}
	dir_list, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalln(err)
	}

	NumPic := len(dir_list)
	if NumPic == 0 {
		log.Fatalf("the pic folder [%s] is empty!", path)
	}

	PicName := dir_list[RandInt(0, NumPic)].Name()
	pwd, _ := os.Getwd()
	PicName = pwd + Separator() + path + Separator() + PicName
	log.Printf("PicName is %s\n", PicName)

	return PicName

}

func main() {
	//BotName := Config.BotName

	setting := tb.Settings{
		Token:  Config.Secure.BotToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	}
	bot, err := tb.NewBot(setting)

	if err != nil {
		log.Fatalln(err)
		return
	}

	bot.Handle("/hi", func(m *tb.Message) {
		bot.Send(m.Chat, "Hello, "+m.Sender.FirstName+"!")
	})

	bot.Handle("/pic", func(m *tb.Message) {
		pic := GenRandomPicFromPath(Config.PicFolder)

		p := &tb.Photo{File: tb.FromDisk(pic)}
		bot.Send(m.Chat, p)
	})

	bot.Handle("/time", func(m *tb.Message) {
		PicUrl := GetBiJinPhoto()
		p := &tb.Photo{File: tb.FromURL(PicUrl)}
		bot.Send(m.Chat, p)
	})

	bot.Start()

}
