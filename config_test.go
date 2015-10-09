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

func CheckInternalValid(t *testing.T, public Config) {
	private := public.internal()
	root, has_root := private.Loggers[""]
	db, has_db := private.Loggers["db"]
	http, has_http := private.Loggers["http"]
	http_request, has_http_request := private.Loggers["http.request"]

	assert.True(t, has_root)
	assert.True(t, has_db)
	assert.True(t, has_http)
	assert.True(t, has_http_request)
	assert.Equal(t, INFO, root.Level)
	assert.Equal(t, INFO, db.Level)
	assert.Equal(t, ERROR, http_request.Level)
	assert.EqualValues(t, root.Pathes, map[string]bool{"/var/log/sample.log": true})
	assert.EqualValues(t, db.Pathes, map[string]bool{"/var/log/sample.log": true})
	assert.EqualValues(t, http.Pathes, map[string]bool{"/var/log/sample.log": true})
	assert.EqualValues(t, http_request.Pathes,
		map[string]bool{
			"/var/log/sample.log": true,
			"/var/log/http.log":   true})
}

func CheckValidConfig(t *testing.T, err error, public Config) {
	assert.NoError(t, err)
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

	CheckInternalValid(t, public)
}
