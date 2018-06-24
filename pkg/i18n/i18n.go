// Copyright 2016-2018 The Aiicy Team.
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
package i18n

import (
	"fmt"
	"path"

	"github.com/Aiicy/AiicyBot/pkg/common"
	"github.com/Unknwon/i18n"
	"golang.org/x/text/language"
)

// Options represents a struct for specifying configuration options for the i18n middleware.
type Options struct {
	// Directory to load locale files. Default is "conf/locale"
	Directory string
	// File stores actual data of locale files. Used for in-memory purpose.
	Files map[string][]byte
	// Custom directory to overload locale files. Default is "custom/conf/locale"
	CustomDirectory string
	// Langauges that will be supported, order is meaningful.
	Langs []string
	// Human friendly names corresponding to Langs list.
	Names []string
	// Default language locale, leave empty to remain unset.
	DefaultLang string
	// Locale file naming style. Default is "locale_%s.ini".
	Format string
	// Name that maps into template variable. Default is "i18n".
	TmplName string
	// Configuration section name. Default is "i18n".
	Section string
}

// initLocales initializes language type list
func initLocales(opt Options) language.Matcher {
	tags := make([]language.Tag, len(opt.Langs))
	for i, lang := range opt.Langs {
		tags[i] = language.Raw.Make(lang)
		fname := fmt.Sprintf(opt.Format, lang)
		// Append custom locale file.
		custom := []interface{}{}
		customPath := path.Join(opt.CustomDirectory, fname)
		if common.IsFile(customPath) {
			custom = append(custom, customPath)
		}

		var locale interface{}
		if data, ok := opt.Files[fname]; ok {
			locale = data
		} else {
			locale = path.Join(opt.Directory, fname)
		}

		err := i18n.SetMessageWithDesc(lang, opt.Names[i], locale, custom...)
		if err != nil && err != i18n.ErrLangAlreadyExist {
			panic(fmt.Errorf("fail to set message file(%s): %v", lang, err))
		}
	}
	return language.NewMatcher(tags)
}

// A Locale describles the information of localization.
type Locale struct {
	i18n.Locale
}

// Language returns language current locale represents.
func (l Locale) Language() string {
	return l.Lang
}

func prepareOptions(options []Options) Options {
	var opt Options
	if len(options) > 0 {
		opt = options[0]
	}

	if len(opt.Section) == 0 {
		opt.Section = "i18n"
	}
	i18n.SetDefaultLang(opt.DefaultLang)

	if len(opt.Directory) == 0 {
		opt.Directory = "locales"
	}
	if len(opt.CustomDirectory) == 0 {
		opt.CustomDirectory = "custom/conf/locale"
	}
	if len(opt.Format) == 0 {
		opt.Format = "locale_%s.ini"
	}
	if len(opt.TmplName) == 0 {
		opt.TmplName = "i18n"
	}

	return opt
}

type LangType struct {
	Lang, Name string
}

var isNeedRedir bool

// I18n is a middleware provides localization layer for your application.
// Paramenter langs must be in the form of "en-US", "zh-CN", etc.
// Otherwise it may not recognize browser input.
func I18n(options ...Options) {
	opt := prepareOptions(options)
	initLocales(opt)
}

func IsExist(Langs []string, lang string) bool {
	for _, v := range Langs {
		if v == lang {
			return true
		}
	}
	return false
}
