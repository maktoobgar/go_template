package app

import (
	"log"

	"github.com/maktoobgar/bookstore/internal/config"
	"github.com/maktoobgar/bookstore/internal/contract"
	"github.com/maktoobgar/bookstore/pkg/translator"
	"golang.org/x/text/language"
)

var (
	languages = []language.Tag{language.English, language.Persian}
)

// Initialize for translation in different languages
func initializeTranslation() (contract.Translator, error) {
	Translator, err := translator.New(languages[0], languages[1:]...)
	if err != nil {
		log.Fatalln(err)
	}
	return Translator, nil
}

func Run(cfg *config.Config) error {
	Translator, _ := initializeTranslation()
	_ = Translator
	return nil
}
