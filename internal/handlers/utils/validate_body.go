package utils

import (
	"github.com/golodash/galidator"
	"github.com/maktoobgar/go_template/pkg/errors"
	"github.com/maktoobgar/go_template/pkg/translator"
)

func ValidateBody(data any, validator galidator.Validator, translate translator.TranslatorFunc) {
	if errs := validator.Validate(data, galidator.Translator(translate)); errs != nil {
		panic(errors.New(errors.InvalidStatus, errors.Resend, translate("BodyNotProvidedProperly"), errs))
	}
}
