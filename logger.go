package logging

import (
	"strings"
	"sync"
)

type Path struct {
	data []string
}

type Logger struct {
	// path to logger, i.e. for logger "http.request" it would be ["http", "request"]
	path Path
	// children loggers
	children map[string]*Logger
	// files attached to logger
	file FileArray
	// logger level
	level Level
}

var root *Logger = nil
var mutex sync.Mutex

func init() {
	root = &Logger{
		Path{},
		make(map[string]*Logger),
		FileArray{},
		INFO,
	}
}

func (self *Path) String() string {
	return strings.Join(self.data, ".")
}

func ParsePath(fullPath string) Path {
	result := make([]string, 0)
	for _, path := range strings.Split(fullPath, ".") {
		if len(path) > 0 {
			result = append(result, path)
		}
	}
	return Path{result}
}

func (self *Path) Iter() []string {
	return self.data
}

func (self *Path) GetChildren(name string) Path {
	result := make([]string, len(self.data)+1)
	copy(result, self.data)
	result[len(self.data)] = name
	return Path{result}
}

func (self *Logger) GetChildren(name string) *Logger {
	result, ok := self.children[name]
	if ok {
		return result
	}
	logger := &Logger{
		self.path.GetChildren(name),
		make(map[string]*Logger),
		FileArray{},
		INFO,
	}
	self.children[name] = logger
	return logger
}

func GetLogger(fullPath string) *Logger {
	current := root
	parsedFullPath := ParsePath(fullPath)
	for _, path := range parsedFullPath.Iter() {
		current = root.GetChildren(path)
	}
	return current
}
