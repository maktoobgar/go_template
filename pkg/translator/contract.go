package translator

type Translator interface {
	TranslateFunction(language string) func(string) string
}
