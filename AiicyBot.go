//
// Copyright (c) 2016-2018 The Aiicy Team
// Licensed under The MIT License (MIT)
// See: LICENSE
//
// +build !windows

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	cusSet "github.com/Aiicy/AiicyBot/pkg/setting"
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
		Token:  cusSet.TgConf.Secure.BotToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	}
	bot, err := tb.NewBot(setting)

	if err != nil {
		log.Fatalln(err)
		return
	}

	bot.Handle(tb.OnText, func(m *tb.Message) {
		/*
			text := fmt.Sprintf("机器人名字 @%s\n", conf.BotName)
			text = text + "Git 服务器地址: [https://git.aiicy.com/](https://git.aiicy.com/)\n"
			text = text + "请点击连接添加我们的机器人: [https://t.me/AiicyBot](https://t.me/AiicyBot)\n"
			text = text + "获取使用方法，请输入 /help\n"
		*/
		ProcessMessage(bot, m)

	})
	bot.Handle("/hi", func(m *tb.Message) {
		bot.Send(m.Chat, Tr("hello", m.Sender.FirstName))
	})

	bot.Handle("/ping", func(m *tb.Message) {
		bot.Send(m.Chat, "pong")
	})

	bot.Handle("/pic", func(m *tb.Message) {
		pic := GenRandomPicFromPath(cusSet.TgConf.PicFolder)

		p := &tb.Photo{File: tb.FromDisk(pic)}
		bot.Notify(m.Chat, tb.UploadingPhoto)
		bot.Send(m.Chat, p)
	})

	bot.Handle("/time", func(m *tb.Message) {
		PicUrl := GetBiJinPhoto()
		p := &tb.Photo{File: tb.FromURL(PicUrl)}
		bot.Notify(m.Chat, tb.UploadingPhoto)
		bot.Send(m.Chat, p)
	})

	bot.Handle("/setlang", func(m *tb.Message) {
		if "zh_CN" == m.Payload || "简体中文" == m.Payload {
			cusSet.CurLang = "zh-CN"
		} else {
			cusSet.CurLang = "en-US"
		}
		bot.Send(m.Chat, Tr("change_lang"))
	})

	bot.Handle("/help", func(m *tb.Message) {
		text := fmt.Sprintf(Tr("help.title", cusSet.TgConf.BotName) + "\n")
		text = text + "/hi -- " + Tr("help.hi") + " \n"
		text = text + "/pic -- " + Tr("help.pic") + " \n"
		text = text + fmt.Sprintf("/time -- "+Tr("help.time", cusSet.TgConf.BijinTZ)+"\n")
		text = text + "/ping -- " + Tr("help.ping") + " \n"
		text = text + "/help -- " + Tr("help.help") + " \n"
		bot.Notify(m.Chat, tb.Typing)
		bot.Send(m.Chat, text)

	})

	bot.Start()

}
