package main

import (
	"github.com/jinzhu/configor"
)

type Config struct {
	BotName   string `default:"AiicyBot"`
	PicFolder string `default:"pics" env:" BOT_PIC_F"`
	BijinTZ   string `required:"true" env:"BIJIN_TIMEZONE"`
	Secure    struct {
		BotToken string `required:"true" env:"BOT_TOKEN"`
	}
}

var conf Config

func init() {
	configor.Load(&conf, "config.yaml")
}
