package main

import (
	"github.com/jinzhu/configor"
)

var Config = struct {
	BotName   string `default:"AiicyBot"`
	PicFolder string `default:"pics"`

	Secure struct {
		BotToken string `required:"true" env:"BOT_TOKEN"`
	}
}{}

func init() {
	configor.Load(&Config, "config.yaml")
}
