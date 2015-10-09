package logging

import (
	"fmt"
	"os"
)

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

func newEmptyConfig() config {
	return config{make(map[string]*loggerConfig)}
}

func newConfig(external Config) config {
	// verify names
	for nameString := range external.InvalidNames() {
		panic(fmt.Sprintf("logger hierarchical name %s normalized to %s, improper configuration",
			nameString, ParseName(nameString).String()))
	}

	result := newEmptyConfig()

	// build loggers
	for nameString, source := range external.Loggers {
		current := result.GetLogger(nameString)
		if err := current.Level.Parse(source.Level); err != nil {
			panic(err)
		}
		if len(source.File) > 0 {
			path := joinPathes(external.Directory, source.File)
			current.Pathes[path] = true
		}
	}

	result.propagatePathes()

	return result
}

func newLoggerConfig(name Name) *loggerConfig {
	pathes := make(map[string]bool)
	return &loggerConfig{name, DEFAULT, pathes}
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
