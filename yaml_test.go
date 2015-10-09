package logging

import (
	"testing"
)

func TestYAML_Positive(t *testing.T) {
	ParseCheckConfig_Positive(t, "yaml", validYAML)
}

func TestYAML_Negative(t *testing.T) {
	ParseCheckConfig_Negative(t, "yaml", invalidYAML)
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
var invalidYAML = `
directory: /var/log
loggers:
  "":
    file: sample.log
    level: INFO
  .db:
  http..request:
    file: http.log
    level: ERROR
`
