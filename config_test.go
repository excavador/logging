package logging

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"testing"
)

func (self Config) Print() {
	result, err := json.Marshal(self)
	fmt.Printf("error: %v data:\n%v\n", err, string(result))
	result, err = yaml.Marshal(self)
	fmt.Printf("error: %v data:\n%v\n", err, string(result))
}

func CheckValidConfig(t *testing.T, public Config) {
	assert.Equal(t, "/var/log", public.Directory)

	public.Print()

	root, has_root := public.Loggers[""]
	db, has_db := public.Loggers["db"]
	http_request, has_http_request := public.Loggers["http.request"]
	_, has_http := public.Loggers["http"]

	assert.True(t, has_root)
	assert.True(t, has_db)
	assert.True(t, has_http_request)
	assert.False(t, has_http)
	assert.Equal(t, INFO, root.Level)
	assert.Equal(t, INFO, db.Level)
	assert.Equal(t, ERROR, http_request.Level)
	assert.Equal(t, root.File, "sample.log")
	assert.Equal(t, db.File, "")
	assert.Equal(t, http_request.File, "http.log")
}

func TestValidJSON(t *testing.T) {
	var public Config
	json.Unmarshal([]byte(validJSON), &public)
	CheckValidConfig(t, public)
}

func TestValidYAML(t *testing.T) {
	var public Config
	yaml.Unmarshal([]byte(validYAML), &public)
	CheckValidConfig(t, public)
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

var validYAML = `
directory: /var/log
loggers:
  "":
    file: sample.log
    level: INFO
  db:
  http.request:
    file: http.log
    level: ERROR
`

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
