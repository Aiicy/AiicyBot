package main

import (
	"fmt"

	"github.com/nlopes/slack"
	"os"
)

func main() {
	var token string = os.Getenv("slacktoken")
	channel := []string{"#general"}
	api := slack.New(token)
	params := slack.FileUploadParameters{
		Title:    "Batman Example",
		Filetype: "auto",
		File:     "example.txt", // *.jpg| *.jpeg|*.png
		Channels: channel,

		//Content:  "Nan Nan Nan Nan Nan Nan Nan Nan Batman",
	}
	file, err := api.UploadFile(params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URLPrivateDownload)

	/*
		err = api.DeleteFile(file.ID)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("File %s deleted successfully.\n", file.Name)
	*/
}
