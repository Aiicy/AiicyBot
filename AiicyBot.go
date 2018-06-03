package main

import (
	"log"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v1"
)

func main() {
	BotName := "AiicyBot"
	bot, err := tb.NewBot(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatalln(err)
	}

	messages := make(chan tb.Message, 100)
	bot.Listen(messages, 10*time.Second)

	pic, err := tb.NewFile("test.jpg")
	if err != nil {
		log.Fatalln(err)
	}
	photo := tb.Photo{File: pic}
	picbot := "/pic@" + BotName
	for message := range messages {
		if message.Text == "/hi" {
			bot.SendMessage(message.Chat,
				"Hello, "+message.Sender.FirstName+"!", nil)
		}

		if message.Text == "/pic" || message.Text == picbot {
			bot.SendPhoto(message.Chat, &photo, nil)

		}

	}
}
