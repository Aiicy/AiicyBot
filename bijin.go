package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	cfg, err := LoadConfig("config.yaml")
	if err != nil {
		loc, _ := time.LoadLocation("Local")
		return loc
	}
	tz := cfg.GetBiJinTimeZone()

	l, err := time.LoadLocation(tz)
	if err != nil {
		log.Printf("Error loading location: %v", err)
	}

	return l
}
func GetBiJinPhoto(region string) (link string, outstr string) {
	rand.Seed(time.Now().UTC().UnixNano())
	t := time.Now().In(LoadLocation())
	hours := fmt.Sprintf("%02d", t.Hour())
	minutes := fmt.Sprintf("%02d", t.Minute())
	region, link, profileLink := getLink(strings.ToLower(strings.TrimSpace(region)))
	prof := getProfile(profileLink, hours, minutes)
	link = fmt.Sprintf("%s%s%s.jpg", link, hours, minutes)
	outstr = fmt.Sprintf("Here's your <%s%s%s.jpg|%s:%s 美人 (%s)>\n%s", link, hours, minutes, hours, minutes, strings.ToTitle(region), prof)

	return link, outstr
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

func getLink(region string) (string, string, string) {
	var r string
	var i string
	if i, ok := prefixes[region]; ok {
		return region, fmt.Sprintf(picTempl, i), fmt.Sprintf(infoTempl, i)
	}
	for r, i = range prefixes {
		break
	}
	return r, fmt.Sprintf(picTempl, i), fmt.Sprintf(infoTempl, i)
}
