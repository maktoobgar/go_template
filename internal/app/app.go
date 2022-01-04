package app

import (
	"github.com/gotorn/core/internal/config"
	"github.com/gotorn/core/pkg/logging"
	"github.com/gotorn/core/pkg/translator"
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
