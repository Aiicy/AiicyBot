//
// Copyright (c) 2016-2018 The Aiicy Team
// Licensed under The MIT License (MIT)
// See: LICENSE
//

package main

//go:generate go-bindata -pkg main -o bindata.go locales/...

import (
	"github.com/Aiicy/AiicyBot/pkg/i18n"
	"github.com/Aiicy/AiicyBot/pkg/setting"
	in "github.com/Unknwon/i18n"
)

func InitI18n() {
	localeNames, err := AssetDir("locales")
	if err != nil {
		log.Errorf("Fail to list locale files: %v\n", err)
	}
	//I18n Local files
	localFiles := make(map[string][]byte)
	for _, name := range localeNames {
		localFiles[name] = MustAsset("locales/" + name)
	}
	i18n.I18n(i18n.Options{
		Files:       localFiles,
		Langs:       setting.Langs,
		Names:       setting.Names,
		DefaultLang: setting.CurLang,
	})
}

func Tr(s string, args ...interface{}) string {
	return in.Tr(setting.CurLang, s, args)
}
