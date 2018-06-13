package main

import (
	"reflect"

	apiai "github.com/mlabouardy/dialogflow-go-client/models"
	"github.com/mlabouardy/moviedb"
	tb "gopkg.in/tucnak/telebot.v2"
)

func ProcessMessage(bot *tb.Bot, m *tb.Message) {

	var userQuery = m.Text
	var dialogFlowResponse = GetResponse(userQuery, conf.DialogFlow.Token, conf.DialogFlow.Lang)

	if !reflect.DeepEqual(dialogFlowResponse.Metadata, apiai.Metadata{}) && dialogFlowResponse.Metadata.IntentName == "shows" {
		var showType = dialogFlowResponse.Parameters["show-type"]
		db := moviedb.NewMovieDB()

		var shows []moviedb.Show

		if showType == "movie" {
			shows = db.GetNowPlayingMovies()
		} else {
			shows = db.GetAiringTodayShows()
		}

		p := &tb.Photo{File: tb.FromURL(shows[0].Cover), Caption: shows[0].Title}
		bot.Notify(m.Chat, tb.UploadingPhoto)
		bot.Send(m.Chat, p)
	} else {

		text := dialogFlowResponse.Fulfillment.Speech
		bot.Send(m.Chat, text)

	}

}
