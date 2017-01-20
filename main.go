package main

import (
	"github.com/urfave/cli"
	"os"
	"time"
)

func actionStartSlack(c *cli.Context) error {
	token := GetToken("config.yaml")
	channelID = GetChannelId("config.yaml")
	slackRun(token, channelID)

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "aiicySlackBot"
	app.Usage = "Slack bot to get the stock infomation"
	app.Version = "0.5.0"
	app.Compiled = time.Now()
	app.Copyright = "Copyright (c) 2017 sndnvaps<admin@sndnvaps.com>"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "sndnvaps",
			Email: "admin@sndnvaps.com",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "start",
			Usage:  "start slack bot",
			Action: actionStartSlack,
		},
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}

}
