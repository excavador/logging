package logging

import (
	"fmt"
	"sync"
)

type Logger struct {
	// path to logger, i.e. for logger "http.request" it would be ["http", "request"]
	path name
	// children loggers
	children map[string]*Logger
	// files attached to logger
	file fileArray
	// logger level
	level Level
}

var root *Logger = nil
var mutex sync.Mutex

func init() {
	root = &Logger{
		name{},
		make(map[string]*Logger),
		fileArray{},
		INFO,
	}
}

func (self *Logger) GetLogger(name string) *Logger {
	mutex.Lock()
	defer mutex.Unlock()
	result, ok := self.children[name]
	if ok {
		return result
	}
	logger := &Logger{
		self.path.GetChild(name),
		make(map[string]*Logger),
		self.file,
		INFO,
	}
	self.children[name] = logger
	return logger
}

func (self *Logger) String() string {
	mutex.Lock()
	defer mutex.Unlock()
	var children []string
	for name, _ := range self.children {
		children = append(children, name)
	}
	return fmt.Sprintf("Logger{path: '%v' children: %v file: %v level %v}",
		self.path, children, self.file, self.level)
}

func GetLogger(name string) *Logger {
	current := root
	parsedName := parseName(name)
	for _, name := range parsedName.data {
		current = current.GetLogger(name)
	}
	return current
}

func Configure(config Config) {
	for _, loggerConfig := range config.Loggers {
		fmt.Printf("%v\n", loggerConfig)
	}
}
