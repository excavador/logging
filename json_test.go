package logging

import (
	"testing"
)

func TestJSON_Positive(t *testing.T) {
	ParseCheckConfig_Positive(t, "json", validJSON)
}

func TestJSON_Negative(t *testing.T) {
	ParseCheckConfig_Negative(t, "json", invalidJSON)
}

var validJSON = `
{
	"directory": "/var/log",
	"loggers": {
		"": {
			"file": "sample.log",
			"level": "INFO"
		},
		"db": {
		},
		"http.request": {
			"file": "http.log",
			"level": "ERROR"
		}
	}
}`

var invalidJSON = `
{
	"directory": "/var/log",
	"loggers": {
		"": {
			"file": "sample.log",
			"level": "INFO"
		},
		".db": {
		},
		"http..request": {
			"file": "http.log",
			"level": "ERROR"
		}
	}
}`
