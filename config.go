package logging

import (
	"fmt"
	"os"
)

// Every logger has unique hierarchical name, like "a.b.c", or "http.request"
type LoggerConfig struct {
	// File name for write data.
	File string `json:"file,omitempty",yaml:"file"`
	// Level
	Level Level `json:"level",yaml:"level"`
}

type Config struct {
	// Directory for place log files
	Directory string `json:"directory",yaml:"directory"`
	// Loggers configuration
	Loggers map[string]LoggerConfig `json:"loggers,omitempty",yaml:"loggers,omitempty"`
}

type loggerConfig struct {
	// Logger name
	Name Name
	// Logger level
	Level Level
	// Pathes of attached files
	Pathes map[string]bool
}

type config struct {
	Loggers map[string]*loggerConfig
}

func (self LoggerConfig) MarshalYAML() (interface{}, error) {
	result := make(map[string]string)
	if len(self.File) > 0 {
		result["file"] = self.File
	}
	result["level"] = self.Level.String()
	return result, nil
}

func newConfig() config {
	return config{make(map[string]*loggerConfig)}
}

func newLoggerConfig(name Name) *loggerConfig {
	pathes := make(map[string]bool)
	return &loggerConfig{name, DEFAULT, pathes}
}

func (self *Config) InvalidNames() chan string {
	result := make(chan string)
	go func() {
		for nameString, _ := range self.Loggers {
			name := ParseName(nameString)
			if name.String() != nameString {
				result <- nameString
			}
		}
		close(result)
	}()
	return result
}

func (self *Config) internal() config {
	// verify names
	for nameString := range self.InvalidNames() {
		panic(fmt.Sprintf("logger hierarchical name %s normalized to %s, improper configuration",
			nameString, ParseName(nameString).String()))
	}

	result := newConfig()

	// build loggers
	for nameString, source := range self.Loggers {
		current := result.GetLogger(nameString)
		current.Level = source.Level
		if len(source.File) > 0 {
			path := joinPathes(self.Directory, source.File)
			current.Pathes[path] = true
		}
	}

	result.propagatePathes()

	return result
}

func (self *config) GetLogger(nameString string) *loggerConfig {
	if result, found := self.Loggers[nameString]; found {
		return result
	}
	name := ParseName(nameString)
	for parentName := range name.Parents() {
		parentNameString := parentName.String()
		if _, found := self.Loggers[parentNameString]; !found {
			self.Loggers[parentNameString] = newLoggerConfig(parentName)
		}
	}
	result := newLoggerConfig(name)
	self.Loggers[nameString] = result
	return result
}

func (self *config) propagatePathes() {
	for _, logger := range self.Loggers {
		for parentName := range logger.Name.Parents() {
			parentNameString := parentName.String()
			parentPathes := self.Loggers[parentNameString].Pathes
			for path, _ := range parentPathes {
				logger.Pathes[path] = true
			}
		}
	}
}

func joinPathes(directory string, filename string) string {
	return fmt.Sprintf("%s%c%s", directory, os.PathSeparator, filename)
}
