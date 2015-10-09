package logging

import (
	"testing"
)

func TestValidYAML(t *testing.T) {
	ParseAndCheckValidConfig(t, "yaml", validYAML)
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
