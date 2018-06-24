package main

//go:generate go-bindata -pkg main -o bindata.go locales/...

import (
	"github.com/Unknwon/i18n"
)

func Tr(curLang string, s string, args ...interface{}) string {
	return i18n.Tr(curLang, s, args)
}
