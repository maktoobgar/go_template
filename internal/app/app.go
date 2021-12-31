package app

import (
	"github.com/maktoobgar/bookstore/internal/config"
	"github.com/maktoobgar/bookstore/pkg/logging"
	"github.com/maktoobgar/bookstore/pkg/translator"
	"golang.org/x/text/language"
)

var (
	languages = []language.Tag{language.English, language.Persian}
)

func Run(cfg *config.Config) error {
	// Translator initialization
	translator, err := translator.New(cfg.Translator.Path, languages[0], languages[1:]...)
	if err != nil {
		return err
	}
	_ = translator

	// Logger initialization
	k := cfg.Logging
	l := logging.Option(k)
	logger, err := logging.New(&l)
	if err != nil {
		return err
	}
	_ = logger

	return nil
}
