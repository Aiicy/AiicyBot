// +build !windows

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/kardianos/osext"
	tb "gopkg.in/tucnak/telebot.v2"
)

func GenRandomPicFromPath(path string) string {
	pwd, err := osext.ExecutableFolder()
	if err != nil {
		log.Error(err)
	}

	log.Debugf("current path is %s", pwd)
	log.Debugf("Pic folder is %s", path)
	path = pwd + "/" + path
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

	PicName = path + "/" + PicName
	log.Infof("PicName is %s", PicName)

	return PicName

}

func StartBot() {
	//BotName := Config.BotName

	setting := tb.Settings{
		Token:  conf.Secure.BotToken,
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
		pic := GenRandomPicFromPath(conf.PicFolder)

		p := &tb.Photo{File: tb.FromDisk(pic)}
		bot.Send(m.Chat, p)
	})

	bot.Handle("/time", func(m *tb.Message) {
		PicUrl := GetBiJinPhoto()
		p := &tb.Photo{File: tb.FromURL(PicUrl)}
		bot.Send(m.Chat, p)
	})

	bot.Handle("/help", func(m *tb.Message) {
		text := fmt.Sprintf("%s support commands\n", conf.BotName)
		text = text + "/hi -- just for test the bot online or not\n"
		text = text + "/pic -- Get the random mm pic\n"
		text = text + fmt.Sprintf("/time -- Show the time of Location:[%s] by mm\n", conf.BijinTZ)
		text = text + "/help -- show the above info\n"
		bot.Send(m.Chat, text)

	})

	bot.Start()

}
