package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"os"
	"github.com/olebedev/config"
)

const Magnet_re = `/magnet[^ <'\"]*/i`

func findMagnets(docs string) []string {
	reg := regexp.MustCompilePOSIX(Magnet_re)
	match := reg.FindAllString(docs, -1)
	return match
}
type slackCFG struct {
	cfg *config.Config
}

func LoadConfig(filename string) (slackCFG, error) {
	cfg, err := config.ParseYamlFile(filename)
	if err != nil {
		fmt.Printf("Cannot parse the config file: %s\n", filename)
		return slackCFG{cfg : &config.Config{}}, err
	}
	return slackCFG{cfg: cfg}, nil
}

func (self slackCFG) GetBotFatherId() string {
	botfather := self.cfg.UString("botfather", "")
	if botfather == "" {
		fmt.Printf("Get botfather from env\n")
		botfather = os.Getenv("botfather")
	}

	return botfather
}

func (self slackCFG) GetToken() string {
	token := self.cfg.UString("slacktoken", "")
	if token == "" {
		fmt.Printf("Get slacktoken from env\n")
		token = os.Getenv("slacktoken")
	}
	return token
}

func (self slackCFG)GetChannelId() string {
	channelId := self.cfg.UString("ReportChannel", "")
	if channelId == "" {
		fmt.Printf("Get ReportChannel from env\n")
		channelId = os.Getenv("ReportChannel")
	}
	return channelId
}

//export BIJIN_TIMEZONE=Asia/Tokyo
func (self slackCFG) GetBiJinTimeZone() string {
	timezone := self.cfg.UString("BIJIN_TIMEZONE", "")
	if timezone == "" {
		fmt.Printf("Get BIJIN_TIMEZONE from env\n")
		timezone = os.Getenv("BIJIN_TIMEZONE")
	}
	return timezone
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

//just coomment it , need more test
func getMagnet(CarNo string) string {
	CarNo = strings.ToUpper(CarNo)
	outstr := ""
	url := fmt.Sprintf("'http://www.btmeet.org/search/%s", CarNo)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	//fmt.Println(string(body))
	magnets := findMagnets(string(body))
	for k, v := range magnets {
		outstr = outstr + fmt.Sprintf("\n%d:%s\n", k, v)
	}

	return outstr
}
