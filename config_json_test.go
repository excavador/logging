package logging

import (
	"encoding/json"
	"testing"
)

func TestValidJSON(t *testing.T) {
	public := NewConfig()
	err := json.Unmarshal([]byte(validJSON), &public)
	CheckValidConfig(t, err, public)
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
