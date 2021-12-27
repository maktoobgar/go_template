package translator

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

const (
	// The address where all translation files are stored.
	address string = "build/translations/"
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
	// The main translator object that will be used in the whole services
	Translator *TranslatorPack = &TranslatorPack{
		addedLanguages: []string{},
		localizers:     map[string]*i18n.Localizer{},
	}
	// If you wanna add another format to support in your translation files,
	// you have to add your new format's extention to supportedFormats by hand
	// and give unmarshal of that format to `bundle` variable. like what I did
	// in second line of Setup function for toml format.
	supportedFormats = []string{"json", "toml"}
)

// Initialization of the translation.
//
// You can get your language Tag with using "golang.org/x/text/language"
// library like: language.English
func New(defaultLanguage language.Tag, languages ...language.Tag) (*TranslatorPack, error) {
	Translator.bundle = i18n.NewBundle(defaultLanguage)
	Translator.bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	languages = append(languages, defaultLanguage)

	err := loadLanguages(languages...)
	if err != nil {
		return nil, err
	}

	return Translator, nil
}

// Translates to requested language and if only the language
// is not added before with `loadLanguages` or `Setup` functions,
// `localizer not found` error returns.
//
// You can get your language string code with using "golang.org/x/text/language"
// library like: language.English.String()
func (Translator *TranslatorPack) Translate(language string, messageID string) string {
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
func (Translator *TranslatorPack) TranslateFA(messageID string) string {
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
func (Translator *TranslatorPack) TranslateEN(messageID string) string {
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
	if Translator.bundle == nil {
		return fmt.Errorf("please call Setup function first")
	}

	var exit bool
	for _, lang := range languages {
		exit = false
		for _, langString := range Translator.addedLanguages {
			if lang.String() == langString {
				exit = true
			}
		}
		if !exit {
			Translator.addedLanguages = append(Translator.addedLanguages, lang.String())
			var err error = nil
			var atLeastOneExists bool = false
			for _, format := range supportedFormats {
				_, e := Translator.bundle.LoadMessageFile(fmt.Sprintf(address+fileFormat, lang.String(), format))
				if atLeastOneExists {
					err = nil
				} else if e == nil {
					atLeastOneExists = true
				} else {
					err = e
				}

			}
			if err != nil {
				return fmt.Errorf("'%s' language has no file to read, create one like "+fileFormat+" in %s\n\t\t     supported formats: %s", lang.String(), lang.String(), "json", address, supportedFormats)
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
	localizer, ok := Translator.localizers[language]
	if ok {
		return localizer, nil
	}

	return nil, fmt.Errorf("localizer not found")
}

// Creates localizers for translation to different languages.
func loadLocalizers() {
	for _, lang := range Translator.addedLanguages {
		_, ok := Translator.localizers[lang]
		if ok {
			continue
		}
		Translator.localizers[lang] = i18n.NewLocalizer(Translator.bundle, lang)
	}
}
