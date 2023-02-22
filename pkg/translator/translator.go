package translator

import (
	"fmt"
	"io/fs"

	// "github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

const (
	// Format of how file names are specified in translation
	// folder storage. it's like: message.en.json or message.fa.toml
	fileFormat string = "translation.%s.json"
	folderName string = "translations"
)

type TranslatorPack struct {
	bundle         *i18n.Bundle
	addedLanguages []string
	localizers     map[string]*i18n.Localizer
}

var (
	errLocalizerNotFound = fmt.Errorf("localizer not found")
)

// Initialization of the translation.
//
// You can get your language Tag with using "golang.org/x/text/language"
// library like: language.English
func New(fileSystem fs.FS, defaultLanguage language.Tag, languages ...language.Tag) (Translator, error) {
	translator := &TranslatorPack{
		addedLanguages: []string{},
		localizers:     map[string]*i18n.Localizer{},
	}
	translator.bundle = i18n.NewBundle(defaultLanguage)
	// translator.bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	languages = append(languages, defaultLanguage)
	err := translator.loadLanguages(languages...)
	if err != nil {
		return nil, err
	}
	err = translator.loadLocalizers(fileSystem)
	if err != nil {
		return nil, err
	}

	return translator, nil
}

// Returns a function based on passed language tag.
//
// You can get your language string code with using "golang.org/x/text/language"
// library like: language.English.String()
func (translator *TranslatorPack) TranslateFunction(language string) TranslatorFunc {
	localizer, err := translator.returnLocalizer(language[:2])
	if err != nil {
		return func(word string) string { return word }
	}

	return func(word string) string {
		return translator.translateLocal(localizer, &i18n.LocalizeConfig{
			MessageID: word,
		})
	}
}

// Used internally to initially load languages into the translator.
//
// You can get your language Tag with using "golang.org/x/text/language"
// library like: language.English
func (translator *TranslatorPack) loadLanguages(languages ...language.Tag) error {
	if translator.bundle == nil {
		return fmt.Errorf("please call `New` function first")
	}

	for _, lang := range languages {
		translator.addedLanguages = append(translator.addedLanguages, lang.String())
	}

	return nil
}

// Translates to preferred localizer.
//
// No error will be returned and if no translation been found,
// same `MessageID` in `config` variable returns.
//
// You can get your desired `localizer` from `returnLocalizer` function.
func (translator *TranslatorPack) translateLocal(localizer *i18n.Localizer, config *i18n.LocalizeConfig) string {
	msg, err := localizer.Localize(config)
	if err != nil {
		return config.MessageID
	}

	return msg
}

// Returns preferred localizer based on language code you passed.
//
// If there is no localizer with that language, `localizer not found`
// error returns.
//
// You can get your language string code with using "golang.org/x/text/language"
// library like: language.English.String()
func (translator *TranslatorPack) returnLocalizer(language string) (*i18n.Localizer, error) {
	localizer, ok := translator.localizers[language]
	if ok {
		return localizer, nil
	}

	return nil, errLocalizerNotFound
}

// Creates localizers for translation to different languages.
func (translator *TranslatorPack) loadLocalizers(data fs.FS) error {
	for _, lang := range translator.addedLanguages {
		translator.localizers[lang] = i18n.NewLocalizer(translator.bundle, lang)
		_, err := translator.bundle.LoadMessageFileFS(data, fmt.Sprintf("%s/%s", folderName, fmt.Sprintf(fileFormat, lang)))
		if err != nil {
			return err
		}
	}
	return nil
}
