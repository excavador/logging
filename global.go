package logging

import (
	"sync"
)

var global struct {
	mutex  sync.Mutex
	Root   *Logger
	Config config
}

func init() {
	global.Root = &Logger{
		Name{},
		make(map[string]*Logger),
		fileArray{},
		INFO,
	}
	global.Config = newEmptyConfig()
}

func GetLogger(name string) *Logger {
	current := global.Root
	for component := range NewName(name).Components() {
		current = current.GetLogger(component)
	}
	return current
}

func SetConfig(config Config) {
	global.mutex.Lock()
	defer global.mutex.Unlock()
}
