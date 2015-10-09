package logging

import (
	"gopkg.in/yaml.v2"
	"testing"
)

func TestValidYAML(t *testing.T) {
	public := NewConfig()
	err := yaml.Unmarshal([]byte(validYAML), &public)
	CheckValidConfig(t, err, public)
}

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
