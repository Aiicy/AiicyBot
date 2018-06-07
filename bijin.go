package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type profile struct {
	Success bool          `json:"success"`
	Result  profileResult `json:"result"`
	Message string        `json:"message"`
}

type profileResult struct {
	ProfileInfo []profileInfo `json:"profile_info"`
}

type profileInfo struct {
	Title string      `json:"title"`
	Note  string      `json:"note"`
	URL   string      `json:"url"`
	Type  profileType `json:"type"`
}

type profileType int

const infoTempl = "http://www.bijint.com/assets/profile/%s/pc/ja/"
const picTempl = "http://www.bijint.com/assets/pict/%s/pc/"

var prefixes = map[string]string{
	"japan":     "jp",
	"thailand":  "thailand",
	"2012":      "2012jp",
	"2011":      "2011jp",
	"tokyo":     "tokyo",
	"hokkaido":  "hokkaido",
	"sendai":    "sendai",
	"akita":     "akita",
	"gunma":     "gunma",
	"niigata":   "niigata",
	"kanazawa":  "kanazawa",
	"fukui":     "fukui",
	"nagoya":    "nagoya",
	"kyoto":     "kyoto",
	"osaka":     "osaka",
	"kobe":      "kobe",
	"okayama":   "okayama",
	"kagawa":    "kagawa",
	"fukuoka":   "fukuoka",
	"kagoshima": "kagoshima",
	"okinawa":   "okinawa",
	"kumamoto":  "kumamoto",
	"saitama":   "saitama",
	"hiroshima": "hiroshima",
	"chiba":     "chiba",
	"nara":      "nara",
	"yamaguchi": "yamaguchi",
	"nagano":    "nagano",
	"shizuoka":  "shizuoka",
	"miyazaki":  "miyazaki",
	"tottori":   "tottori",
	"iwate":     "iwate",
	"ibaraki":   "ibaraki",
	"tochigi":   "tochigi",
	"taiwan":    "taiwan",
	"hawaii":    "hawaii",
	"seifuku":   "seifuku",
	"megane":    "megane",
	"sara":      "sara",
	"hairstyle": "hairstyle",
	"circuit":   "cc",
	"hanayome":  "hanayome",
	"waseda":    "wasedastyle",
}

func LoadLocation() *time.Location {

	tz := conf.BijinTZ

	l, err := time.LoadLocation(tz)
	if err != nil {
		log.Errorf("Error loading location: %s", err.Error())
	}

	return l
}

//region default = jp

func GetBiJinPhoto() (link string) {
	region := "japan"
	rand.Seed(time.Now().UTC().UnixNano())
	t := time.Now().In(LoadLocation())
	hours := fmt.Sprintf("%02d", t.Hour())
	minutes := fmt.Sprintf("%02d", t.Minute())
	link = getLink(strings.ToLower(strings.TrimSpace(region)))
	link = fmt.Sprintf("%s%s%s.jpg", link, hours, minutes)
	log.Infof("Time 's PicUrl = %s", link)
	return link
}

func (p profile) String() string {
	if !p.Success {
		return ""
	}
	msg := ""
	if p.Message != "" {
		msg += p.Message
	}
	for _, r := range p.Result.ProfileInfo {
		if r.Title == "" || r.Note == "" || r.Note == "-" {
			continue
		}
		if r.URL != "" {
			msg += fmt.Sprintf("\n%s: <%s|%s>", r.Title, r.URL, r.Note)
		} else {
			msg += fmt.Sprintf("\n%s: %s", r.Title, r.Note)
		}
	}
	return msg
}

func getProfile(profileLink, hours, minutes string) profile {
	p := profile{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s%s.json", profileLink, hours, minutes), nil)
	if err != nil {
		return p
	}
	req.Header.Add("Host", "http://www.bijint.com")
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return p
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return p
	}
	err = json.Unmarshal(b, &p)
	if err != nil {
		return p
	}
	return p
}

func getLink(region string) string {
	var i string
	if i, ok := prefixes[region]; ok {
		return fmt.Sprintf(picTempl, i)
	}

	return fmt.Sprintf(picTempl, i)
}
