package logging

import (
	"bytes"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func CheckConfigInternal_Positive(t *testing.T, external Config) {
	internal := newConfig(external)

	root, has_root := internal.Loggers[""]
	db, has_db := internal.Loggers["db"]
	http, has_http := internal.Loggers["http"]
	http_request, has_http_request := internal.Loggers["http.request"]

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

func CheckConfig_Positive(t *testing.T, external Config) {
	assert.Equal(t, "/var/log", external.Directory)

	root, has_root := external.Loggers[""]
	db, has_db := external.Loggers["db"]
	http_request, has_http_request := external.Loggers["http.request"]
	_, has_http := external.Loggers["http"]

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

	CheckConfigInternal_Positive(t, external)
}

func Parse(t *testing.T, configType string, configData string) Config {
	data := bytes.NewBuffer([]byte(configData))

	v := viper.New()
	v.SetConfigType(configType)

	err := v.ReadConfig(data)
	assert.NoError(t, err)

	result := NewConfig()
	err = v.Unmarshal(&result)
	assert.NoError(t, err)

	return result
}

func ParseCheckConfig_Positive(t *testing.T, configType string, configData string) {
	config := Parse(t, configType, configData)
	CheckConfig_Positive(t, config)
}

func ParseCheckConfig_Negative(t *testing.T, configType string, configData string) {
	config := Parse(t, configType, configData)
	actual := make(map[string]bool)
	for name := range config.InvalidNames() {
		actual[name] = true
	}
	assert.EqualValues(t, map[string]bool{".db": true, "http..request": true}, actual)
}
