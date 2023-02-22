package translator

type TranslatorFunc func(string) string

type Translator interface {
	TranslateFunction(language string) TranslatorFunc
}
