package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nlopes/slack"
	"strings"
)
const SLACK_TOKEN string = string("xoxb-130783843286-F9pYPTMBPCxhnAJJwVjaEViZ")
var botID = "N/A" // U2NQSPHHD bender bot userID

func main() {
	api := slack.New(SLACK_TOKEN)
	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
	api.SetDebug(true)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Print("Event Received: ")
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				// Ignore hello

			case *slack.ConnectedEvent:
				botID = ev.Info.User.ID
				//rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", "#github"))

			case *slack.MessageEvent:
				fmt.Printf("Message: %v\n", ev)

			case *slack.PresenceChangeEvent:
				fmt.Printf("Presence Change: %v\n", ev)

			case *slack.LatencyReport:
				fmt.Printf("Current latency: %v\n", ev.Value)

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop

			default:
				if msg.Type == "message" && strings.HasPrefix(msg.Text, "<@"+botID+">") {
// if so try to parse if
					parts := strings.Fields(msg.Text)
					if len(parts) == 3 && parts[1] == "stock" {
						// looks good, get the quote and reply with the result
						stock_mes := getQuote(parts[2])
						stock_info := parts[2]
						params := slack.PostMessageParameters{}
						attachment := slack.Attachment{
							Pretext: "some pretext",
							Text:    "some text",
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
						channelID, timestamp, err := api.PostMessage("CHANNEL_ID", "stock ", params)
						if err != nil {
							fmt.Printf("%s\n", err)
							return
						}
						fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
					}
	
				}
				// NOTE: the Message object is copied, this is intentional
				/*
				} else {
				// huh?
					m.Text = fmt.Sprintf("sorry, that does not compute\n")
					postMessage(ws, m)
				}
				*/
				// Ignore other events..
				// fmt.Printf("Unexpected: %v\n", msg.Data)
			}
		}
	}
}