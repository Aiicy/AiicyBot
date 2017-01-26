package main

import (
	"fmt"
	"github.com/nlopes/slack"
	//"github.com/olebedev/config"
	"log"
	"os"
	"strings"
)

var botID = "N/A" // U2NQSPHHD bender bot userID
var botName = "@aiicybot"
var channelID = "N/A" // C3K9VAK3N
//var botFather = "" // use @botname whoami, to get the user_id of bot_father

//cmd = stock aapl
//cmd = whoami
//cmd = botfather | father | bot_father
//cmd = fuli | magnet
//cmd = old | old_driver | olddriver

func ParseCommand(api *slack.Client, rtm *slack.RTM, cmd []string, userid string, botid string, channelid string) error {
	if len(cmd) == 2 && cmd[1] == "whoami" {
		user, err := api.GetUserInfo(userid)
		if err != nil {
			fmt.Printf("%s\n", err)
			return err
		}
		msg := fmt.Sprintf("USER_ID: %s\nFullname: %s\nEmail: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)

		params := slack.PostMessageParameters{}
		attachment := slack.Attachment{
			Pretext:  "who am i",
			Text:     msg,
			ImageURL: user.Profile.Image192,
		}
		params.Attachments = []slack.Attachment{attachment}
		channelID, timestamp, err := api.PostMessage(channelid, "whoami ", params)
		fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	} else if len(cmd) == 2 && (cmd[1] == "help" || cmd[1] == "h") {

		mesg := fmt.Sprintf("*%s* command opts [val]\n", botName)
		mesg = fmt.Sprintf(mesg + "`help|h`, get this help info\n")
		mesg = fmt.Sprintf(mesg + "command list\n")
		mesg = fmt.Sprintf(mesg + "   `whoami` -- get the user_id\n")
		mesg = fmt.Sprintf(mesg + "   `old|old_driver|olddriver`, who is the old driver\n")
		mesg = fmt.Sprintf(mesg + "   `bot_father|botfather|father`, get the bot father name\n")
		mesg = fmt.Sprintf(mesg + "   `fuli|magnet`  val , get the magnet link\n")

		params := slack.PostMessageParameters{}
		//parse markdown
		//https://daringfireball.net/projects/markdown/
		var markdownIn []string = []string{"text"}

		attachment := slack.Attachment{
			Text:       mesg,
			MarkdownIn: markdownIn,
		}
		params.Attachments = []slack.Attachment{attachment}
		channelID, timestamp, err := api.PostMessage(channelid, "", params)
		fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
		if err != nil {
			fmt.Printf("%s\n", err)
			return err
		}
	} else if len(cmd) == 3 && cmd[1] == "stock" {
		// looks good, get the quote and reply with the result
		stock_mes := getQuote(cmd[2])
		stock_info := cmd[2]
		params := slack.PostMessageParameters{}
		attachment := slack.Attachment{
			Pretext: stock_info,
			Text:    stock_mes,
		}
		params.Attachments = []slack.Attachment{attachment}
		channelID, timestamp, err := api.PostMessage(channelid, "stock ", params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return err
		}
		fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	} else if len(cmd) == 3 && (cmd[1] == "fuli" || cmd[1] == "magnet") {
		// looks good, get the quote and reply with the result

		//magnets := getMagnet(cmd[2])

		magnet_info := cmd[2]
		params := slack.PostMessageParameters{}
		attachment := slack.Attachment{
			Pretext: magnet_info,
			Text:    "test joy",
			//Text:    magnets,
		}
		params.Attachments = []slack.Attachment{attachment}
		channelID, timestamp, err := api.PostMessage(channelid, ":unamused: Magnet ", params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return err
		}
		fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	} else if len(cmd) == 2 && (cmd[1] == "old_driver" || cmd[1] == "old" || cmd[1] == "olddriver") {

		user, err := api.GetUserInfo(userid)
		if err != nil {
			fmt.Printf("%s\n", err)
			return err
		}
		msg := fmt.Sprintf(":point_right: %s is an old driver\n", user.Profile.RealName)

		params := slack.PostMessageParameters{}
		attachment := slack.Attachment{
			//Pretext: magnet_info,
			Text: msg,
			//Text:    magnets,
		}
		params.Attachments = []slack.Attachment{attachment}
		channelID, timestamp, err := api.PostMessage(channelid, "", params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return err
		}
		fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)

	} else if len(cmd) == 2 && (cmd[1] == "botfather" || cmd[1] == "father" || cmd[1] == "bot_father") {

		cfg, err := LoadConfig("config.yaml")
		if err != nil {
			return err
		}

		botfather := cfg.GetBotFatherId() // get botfathe userid from config.yaml
		user, err := api.GetUserInfo(botfather)
		if err != nil {
			fmt.Printf("%s\n", err)
			return err
		}

		msg := fmt.Sprintf(":point_right: my botfather %s build me as a robot \n", user.Profile.RealName)

		params := slack.PostMessageParameters{}
		attachment := slack.Attachment{
			//Pretext: magnet_info,
			Text: msg,
			//Text:    magnets,
		}
		params.Attachments = []slack.Attachment{attachment}
		channelID, timestamp, err := api.PostMessage(channelid, "", params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return err
		}
		fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)

	} else if len(cmd) == 3 && (cmd[1] == "bijin" || cmd[1] == "pic") {

		piclink, msg := GetBiJinPhoto(cmd[2])
		params := slack.PostMessageParameters{}
		attachment := slack.Attachment{
			//Pretext: magnet_info,
			Text:     msg,
			ImageURL: piclink,
		}
		params.Attachments = []slack.Attachment{attachment}
		channelID, timestamp, err := api.PostMessage(channelid, "", params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return err
		}
		fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)

	} else if len(cmd) == 3 && (cmd[1] == "weather" || cmd[1] == "forecast") {
		//input should be 
		//@aiicybot weather `foshan,guangdong,china`
		msg, icon_url  := GetWeatherForecastInfo(cmd[2])
		params := slack.PostMessageParameters{}
		attachment := slack.Attachment{
			//Pretext: magnet_info,
			Text:     msg,
			ImageURL: icon_url,
		}
		params.Attachments = []slack.Attachment{attachment}
		channelID, timestamp, err := api.PostMessage(channelid, "", params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return err
		}
		fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)

	} else if len(cmd) == 2 && (cmd[1] == "hentai" || cmd[1] == "pic" ||
		 cmd[1] == "mm" || cmd[1] == "fuli") {

		cfg, err := LoadConfig("config.yaml")
		if err != nil {
			return err
		}
		image_path := cfg.GetImagePath()
		dir_list := ListPath(image_path)
		params := slack.FileUploadParameters{
			Title:    "sex mm pic",
			Filetype: "auto",
			File:     RandomPicName(dir_list), // *.jpg| *.jpeg|*.png|*.gif
			Channels: []string{channelid},

			//Content:  "Nan Nan Nan Nan Nan Nan Nan Nan Batman",
		}
		file, err := api.UploadFile(params)
		
		if err != nil {
			fmt.Printf("%s\n", err)
			return err
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URLPrivateDownload)

	}
	return nil

}

func slackRun(token string, channelid string) {

	api := slack.New(token)
	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
	api.SetDebug(true)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			//fmt.Print("Event Received: ")
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				// Ignore hello

			case *slack.ConnectedEvent:
				botID = ev.Info.User.ID
				//channelId = slack.Msg.Channel
				//rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", "#github"))

			case *slack.MessageEvent:
				callerID := ev.Msg.User
				if ev.Msg.Type == "message" && callerID != botID && ev.Msg.SubType != "message_deleted" &&
					(strings.Contains(ev.Msg.Text, "<@"+botID+">") || strings.HasPrefix(ev.Msg.Channel, "g")) {
					logger.Printf("Message: %+v\n", ev.Msg)
					originalMessage := ev.Msg.Text

					parts := strings.Fields(originalMessage)

					ParseCommand(api, rtm, parts, callerID, botID, channelid)
				}

			case *slack.PresenceChangeEvent:
				fmt.Printf("Presence Change: %v\n", ev)

			case *slack.LatencyReport:
				api.GetUserInfo(botID)

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop

			default:
				// NOTE: the Message object is copied, this is intentional
				// Ignore other events..
				// fmt.Printf("Unexpected: %v\n", msg.Data)
			}
		}
	}

}
