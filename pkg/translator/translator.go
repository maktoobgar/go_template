package translator

import (
	"fmt"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

const (
	// Format of how file names are specified in translation
	// folder storage. it's like: message.en.json or message.fa.toml
	fileFormat string = "translation.%s.%s"
)

type TranslatorPack struct {
	bundle         *i18n.Bundle
	addedLanguages []string
	localizers     map[string]*i18n.Localizer
}

var (
	// The address where all translation files are stored.
	path string = ""
	// The main translator object that will be used in the whole services
	translator *TranslatorPack = &TranslatorPack{
		addedLanguages: []string{},
		localizers:     map[string]*i18n.Localizer{},
	}
	// If you wanna add another format to support in your translation files,
	// you have to add your new format's extention to supportedFormats by hand
	// and give unmarshal of that format to `bundle` variable. like what I did
	// in second line of Setup function for toml format.
	supportedFormats               = []string{"json", "toml"}
	errOperationSystemNotSupported = fmt.Errorf("%s operation system not supported", runtime.GOOS)
	errLocalizerNotFound           = fmt.Errorf("localizer not found")
)

// Initialization of the translation.
//
// You can get your language Tag with using "golang.org/x/text/language"
// library like: language.English
//
// If `address` == "", on linux we will set address to "build/translations" and
// on windows address is "build\translations\"
//
// TODO: Test for windows required
func New(address string, defaultLanguage language.Tag, languages ...language.Tag) (Translator, error) {
	path = address

	if path == "" {
		if runtime.GOOS == "linux" {
			path = "build/translations/"
		} else if runtime.GOOS == "windows" {
			path = "build\\translations\\"
		} else {
			return nil, errOperationSystemNotSupported
		}
	}

	if path[len(path)-1] != '/' || path[len(path)-1] != '\\' {
		if runtime.GOOS == "linux" {
			path += "/"
		} else if runtime.GOOS == "windows" {
			path += "\\"
		} else {
			return nil, errOperationSystemNotSupported
		}
	}

	translator.bundle = i18n.NewBundle(defaultLanguage)
	translator.bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	languages = append(languages, defaultLanguage)

	err := loadLanguages(languages...)
	if err != nil {
		return nil, err
	}

	return translator, nil
}

// Translates to requested language and if only the language
// is not added before with `loadLanguages` or `Setup` functions,
// `localizer not found` error returns.
//
// You can get your language string code with using "golang.org/x/text/language"
// library like: language.English.String()
func (translator *TranslatorPack) Translate(language string, messageID string) string {
	localizer, err := returnLocalizer(language)
	if err != nil {
		return ""
	}

	return translateLocal(localizer, &i18n.LocalizeConfig{
		MessageID: messageID,
		DefaultMessage: &i18n.Message{
			ID: messageID,
		},
	})
}

// Translates to English and if only the language
// is not added before with `loadLanguages` or `Setup` functions,
// `localizer not found` error returns.
//
// You can get your language string code with using "golang.org/x/text/language"
// library like: language.English.String()
func (translator *TranslatorPack) TranslateFA(messageID string) string {
	localizer, err := returnLocalizer(language.Persian.String())
	if err != nil {
		return ""
	}

	return translateLocal(localizer, &i18n.LocalizeConfig{
		MessageID: messageID,
		DefaultMessage: &i18n.Message{
			ID: messageID,
		},
	})
}

// Translates to English and if only the language
// is not added before with `loadLanguages` or `Setup` functions,
// `localizer not found` error returns.
//
// You can get your language string code with using "golang.org/x/text/language"
// library like: language.English.String()
func (translator *TranslatorPack) TranslateEN(messageID string) string {
	localizer, err := returnLocalizer(language.English.String())
	if err != nil {
		return ""
	}

	return translateLocal(localizer, &i18n.LocalizeConfig{
		MessageID: messageID,
		DefaultMessage: &i18n.Message{
			ID: messageID,
		},
	})
}

// Loads translation files and if no file for language determined
// in specified address in `address` variable, return error
//
// Also appends new languages to addedLanguages variable
//
// After that it calls loadLocalizers to create localizers for translation
// to different languages
//
// You can get your language Tag with using "golang.org/x/text/language"
// library like: language.English
func loadLanguages(languages ...language.Tag) error {
	if translator.bundle == nil {
		return fmt.Errorf("please call Setup function first")
	}

	var exit bool
	for _, lang := range languages {
		exit = false
		for _, langString := range translator.addedLanguages {
			if lang.String() == langString {
				exit = true
			}
		}
		if !exit {
			translator.addedLanguages = append(translator.addedLanguages, lang.String())
			var err error = nil
			var atLeastOneExists bool = false
			for _, format := range supportedFormats {
				_, e := translator.bundle.LoadMessageFile(fmt.Sprintf(path+fileFormat, lang.String(), format))
				if atLeastOneExists {
					err = nil
				} else if e == nil {
					atLeastOneExists = true
				} else {
					err = e
				}

			}
			if err != nil {
				return fmt.Errorf("'%s' language has no file to read, create one like "+fileFormat+" in %s\n\t\t     supported formats: %s", lang.String(), lang.String(), "json", path, supportedFormats)
			}
		}
	}
	loadLocalizers()

	return nil
}

// Translates to preferred localizer.
//
// No error will be returned and if no translation been found,
// same `MessageID` in `config` variable returns.
//
// You can get your desired `localizer` from `returnLocalizer` function.
func translateLocal(localizer *i18n.Localizer, config *i18n.LocalizeConfig) string {
	config.DefaultMessage = &i18n.Message{
		ID:    config.MessageID,
		One:   config.MessageID,
		Other: config.MessageID,
	}

	msg, _ := localizer.Localize(config)
	return msg
}

// Returns preferred localizer based on language code you passed.
//
// If there is no localizer with that language, `localizer not found`
// error returns.
//
// You can get your language string code with using "golang.org/x/text/language"
// library like: language.English.String()
func returnLocalizer(language string) (*i18n.Localizer, error) {
	localizer, ok := translator.localizers[language]
	if ok {
		return localizer, nil
	}

	return nil, errLocalizerNotFound
}

// Creates localizers for translation to different languages.
func loadLocalizers() {
	for _, lang := range translator.addedLanguages {
		_, ok := translator.localizers[lang]
		if ok {
			continue
		}
		translator.localizers[lang] = i18n.NewLocalizer(translator.bundle, lang)
	}
}
