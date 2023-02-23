package logging

import "net/http"

type (
	Logger interface {
		Info(message string, r *http.Request, function any, params ...map[string]any)
		Warning(message string, r *http.Request, function any, params ...map[string]any)
		Error(message string, r *http.Request, function any, params ...map[string]any)
		Panic(err any, r *http.Request, stack string, params ...map[string]any)
	}

	Option struct {
		Path, Pattern, MaxAge, RotationTime, RotationSize string
	}
)
