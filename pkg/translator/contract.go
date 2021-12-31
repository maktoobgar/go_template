package translator

const (
	EN = "en"
	FA = "fa"
)

type (
	Translator interface {
		Translate(language string, messageID string) string
		TranslateEN(messageID string) string
		TranslateFA(messageID string) string
	}
)
