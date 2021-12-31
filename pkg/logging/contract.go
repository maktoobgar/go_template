package logging

type (
	Logger interface {
		Info(message string, function interface{}, params map[string]interface{})
		Warning(message string, function interface{}, params map[string]interface{})
		Error(message string, function interface{}, params map[string]interface{})
		Panic(message string, function interface{}, params map[string]interface{})
	}

	Option struct {
		Path, Pattern, MaxAge, RotationTime, RotationSize string
	}
)
