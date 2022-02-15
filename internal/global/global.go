package g

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maktoobgar/go_template/internal/config"
	"github.com/maktoobgar/go_template/pkg/logging"
	"github.com/maktoobgar/go_template/pkg/translator"
)

var CFG *config.Config = nil
var Logger *logging.LogBundle = nil
var Translator *translator.TranslatorPack = nil
var App *fiber.App = nil

func Log() logging.Logger {
	return Logger
}

func Trans() translator.Translator {
	return Translator
}
