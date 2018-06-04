package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v1"
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

func GenRandomPicFromPath(path string) tb.Photo {
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
	PicName = path + Separator() + PicName
	Pic, _ := tb.NewFile(PicName)

	photo := tb.Photo{File: Pic}

	return photo

}

func main() {
	BotName := Config.BotName
	bot, err := tb.NewBot(Config.Secure.BotToken)
	if err != nil {
		log.Fatalln(err)
	}

	messages := make(chan tb.Message, 100)
	bot.Listen(messages, 10*time.Second)

	picbot := "/pic@" + BotName
	hibot := "/hi@" + BotName
	for message := range messages {
		if message.Text == "/hi" || message.Text == hibot {
			bot.SendMessage(message.Chat,
				"Hello, "+message.Sender.FirstName+"!", nil)
		}

		if message.Text == "/pic" || message.Text == picbot {

			photo := GenRandomPicFromPath(Config.PicFolder)

			bot.SendPhoto(message.Chat, &photo, nil)

		}

	}
}
