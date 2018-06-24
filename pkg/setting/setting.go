package setting

import (
	"fmt"

	"github.com/go-ini/ini"
)

var (

	//Global setting objects
	Cfg *ini.File

	// I18n settings
	Langs      []string
	Names      []string
	LangsNames map[string]string
	dateLangs  map[string]string
	CurLang    string
)

func LoadConfig() {
	CurLang = "en-US"
	Cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to parse 'conf/app.ini': %v\n", err)
	}
	Cfg.NameMapper = ini.AllCapsUnderscore

	Langs = Cfg.Section("i18n").Key("LANGS").Strings(",")
	Names = Cfg.Section("i18n").Key("NAMES").Strings(",")
	LangsNames = Cfg.Section("i18n.langsnames").KeysHash()
	dateLangs = Cfg.Section("i18n.datelang").KeysHash()
}
