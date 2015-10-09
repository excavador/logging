package logging

import (
	"bytes"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func CheckInternalValid(t *testing.T, public Config) {
	private := newConfig(public)

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

func CheckValidConfig(t *testing.T, public Config) {
	assert.Equal(t, "/var/log", public.Directory)

	root, has_root := public.Loggers[""]
	db, has_db := public.Loggers["db"]
	http_request, has_http_request := public.Loggers["http.request"]
	_, has_http := public.Loggers["http"]

	assert.True(t, has_root)
	assert.True(t, has_db)
	assert.True(t, has_http_request)
	assert.False(t, has_http)
	assert.Equal(t, "INFO", root.Level)
	assert.Equal(t, "", db.Level)
	assert.Equal(t, "ERROR", http_request.Level)
	assert.Equal(t, root.File, "sample.log")
	assert.Equal(t, db.File, "")
	assert.Equal(t, http_request.File, "http.log")

	CheckInternalValid(t, public)
}

func ParseAndCheckValidConfig(t *testing.T, configType string, configData string) {
	data := bytes.NewBuffer([]byte(configData))

	v := viper.New()
	v.SetConfigType(configType)

	err := v.ReadConfig(data)
	assert.NoError(t, err)

	public := NewConfig()
	err = v.Unmarshal(&public)
	assert.NoError(t, err)

	CheckValidConfig(t, public)
}
