package main

import (
	"encoding/csv"
	"fmt"
	"github.com/nlopes/slack"
	"github.com/olebedev/config"
	"log"
	"net/http"
	"os"
	"strings"
)

var botID = "N/A"     // U2NQSPHHD bender bot userID
var channelID = "N/A" // C3K9VAK3N

func GetToken(filename string) string {
	cfg, err := config.ParseYamlFile(filename)
	if err != nil {
		fmt.Printf("Cannot parse the config file: %s\n", filename)
	}
	token := cfg.UString("slacktoken", "")
	if token == "" {
		fmt.Printf("Get slacktoken from env\n")
		token = os.Getenv("slacktoken")
	}
	/*
		if err != nil {
			fmt.Printf("Cannot get the slacktoken from config file: %s\n", filename)
		} else {
			token = os.Getenv("slacktoken")
		}
	*/
	return token
}

func GetChannelId(filename string) string {
	cfg, err := config.ParseYamlFile(filename)
	if err != nil {
		fmt.Printf("Cannot parse the config file: %s\n", filename)
		return ""
	}
	channelId, err := cfg.String("ReportChannel")
	if err != nil {
		fmt.Printf("Cannot get the slacktoken from config file: %s\n", filename)
		return ""
	}
	return channelId
}

//cmd = stock aapl
//cmd = whoami
func ParseCommand(api *slack.Client, rtm *slack.RTM, cmd []string, userid string, channelid string) error {
	if len(cmd) == 3 && cmd[1] == "stock" {
		// looks good, get the quote and reply with the result
		stock_mes := getQuote(cmd[2])
		stock_info := cmd[2]
		params := slack.PostMessageParameters{}
		attachment := slack.Attachment{
			Pretext: stock_info,
			Text:    stock_mes,
			// Uncomment the following part to send a field too
			/*
				Fields: []slack.AttachmentField{
					slack.AttachmentField{
							Title: "a",
							Value: "no",
							},
						},
			*/
		}
		params.Attachments = []slack.Attachment{attachment}
		channelID, timestamp, err := api.PostMessage(channelid, "stock ", params)
		//channelID, timestamp, err := api.PostMessage("C3K9VAK3N", "stock ", params)
		//channelID, timestamp, err := api.PostMessage("CHANNEL_ID", "stock ", params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return err
		}
		fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	} else if len(cmd) == 2 && cmd[1] == "whoami" {
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
			ImageURL: user.Profile.Image72,
			// Uncomment the following part to send a field too
			/*
				Fields: []slack.AttachmentField{
					slack.AttachmentField{
							Title: "a",
							Value: "no",
							},
						},
			*/
		}
		params.Attachments = []slack.Attachment{attachment}
		channelID, timestamp, err := api.PostMessage(channelid, "whoami ", params)
		fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
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

					ParseCommand(api, rtm, parts, callerID, channelid)
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

// Get the quote via Yahoo. You should replace this method to something
// relevant to your team!
func getQuote(sym string) string {
	sym = strings.ToUpper(sym)
	url := fmt.Sprintf("http://download.finance.yahoo.com/d/quotes.csv?s=%s&f=nsl1op&e=.csv", sym)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	rows, err := csv.NewReader(resp.Body).ReadAll()
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	if len(rows) >= 1 && len(rows[0]) == 5 {
		return fmt.Sprintf("%s (%s) is trading at $%s", rows[0][0], rows[0][1], rows[0][2])
	}
	return fmt.Sprintf("unknown response format (symbol was \"%s\")", sym)
}
