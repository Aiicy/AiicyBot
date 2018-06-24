//
// Copyright (c) 2016-2018 The Aiicy Team
// Licensed under The MIT License (MIT)
// See: LICENSE
//

package main

import (
	"github.com/mlabouardy/dialogflow-go-client"
	apiai "github.com/mlabouardy/dialogflow-go-client/models"
)

func GetResponse(input string, token string, lang string) apiai.Result {
	err, client := dialogflow.NewDialogFlowClient(apiai.Options{
		AccessToken: token,
		ApiLang:     lang,
	})
	if err != nil {
		log.Fatal(err)
	}

	query := apiai.Query{
		Query: input,
	}
	resp, err := client.QueryFindRequest(query)
	if err != nil {
		log.Fatal(err)
	}
	return resp.Result
}
