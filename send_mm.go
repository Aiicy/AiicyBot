package main

import (
	"fmt"
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

