package main

import (
	"fmt"

	"github.com/nlopes/slack"
	"io/ioutil"
	"os"
	"strings"
)

//add *.jpg| *.jpeg|*.png|*.gif to -> []string, as return
func ListPath(path string) []string {
	var dl_out []string
	basePath, _ := os.Getwd() //+ "/" + path
	dir_list, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("read dir error ", err.Error())
	}
	for _, key := range dir_list {
		//fmt.Println(index, "=", key.Name())
		if strings.HasSuffix(key.Name(), ".jpg") ||
			strings.HasSuffix(key.Name(), ".jpeg") ||
			strings.HasSuffix(key.Name(), ".png") ||
			strings.HasSuffix(key.Name(), ".gif") {
			dl_out = append(dl_out, basePath+"/"+path+"/"+key.Name())
		}
	}
	return dl_out
}

func RandomPicName(files []string) string {
	index := RandInt(0, len(files))
	return files[index]
}

func main() {
	var token string = os.Getenv("slacktoken")
	dir_list := ListPath("images")
	channel := []string{"#general"}
	api := slack.New(token)
	params := slack.FileUploadParameters{
		Title:    "Batman Example",
		Filetype: "auto",
		File:     RandomPicName(dir_list), // *.jpg| *.jpeg|*.png|*.gif
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
