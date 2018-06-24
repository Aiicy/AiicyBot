//
// Copyright (c) 2016-2018 The Aiicy Team
// Licensed under The MIT License (MIT)
// See: LICENSE
//

package setting

import (
	"fmt"

	"github.com/go-ini/ini"
	"github.com/jinzhu/configor"
)

type Config struct {
	BotName   string `default:"AiicyBot"`
	PicFolder string `default:"pics" env:" BOT_PIC_F"`
	BijinTZ   string `required:"true" env:"BIJIN_TIMEZONE"`
	Secure    struct {
		BotToken string `required:"true" env:"BOT_TOKEN"`
	}
	DialogFlow struct {
		Token string `required:"true" env:"DialogFlow_Token"`
		Lang  string `default:"ZH_CN" env:"DialogFlow_Lang"`
	}
}

var (

	//Global setting objects
	Cfg    *ini.File
	TgConf Config

	// I18n settings
	Langs      []string
	Names      []string
	LangsNames map[string]string
	dateLangs  map[string]string
	CurLang    string
)

func init() {
	configor.Load(&TgConf, "config.yaml")
}

func LoadConfig() {
	CurLang = "en-US"
	Cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to parse 'conf/app.ini': %v\n", err)
	}
	Cfg.NameMapper = ini.AllCapsUnderscore

	Langs = Cfg.Section("i18n").Key("LANGS").Strings(",")
	Names = Cfg.Section("i18n").Key("NAMES").Strings(",")
	LangsNames = Cfg.Section("i18n.langsnames").KeysHash()
	dateLangs = Cfg.Section("i18n.datelang").KeysHash()
}
