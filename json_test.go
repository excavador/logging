package logging

import (
	"testing"
)

func TestValidJSON(t *testing.T) {
	ParseAndCheckValidConfig(t, "json", validJSON)
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
			"file": "sample.log"
		},
		".db": {
			"file": "db.log"
		},
		"http..request": {
			"file": "http.log",
			"level": "DEBUG"
		}
	}
}`
