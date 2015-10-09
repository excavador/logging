package logging

import (
	"strings"
)

// logical path of logger, i.e. for logger "http.request" it would be ["http", "request"]
type path struct {
	data []string
}

func (self path) String() string {
	return strings.Join(self.data, ".")
}

func parsePath(fullPath string) path {
	result := make([]string, 0)
	for _, path := range strings.Split(fullPath, ".") {
		if len(path) > 0 {
			result = append(result, path)
		}
	}
	return path{result}
}

func (self *path) Iter() []string {
	return self.data
}

func (self *path) GetChild(name string) path {
	length := len(self.data)
	result := make([]string, length+1)
	copy(result, self.data)
	result[length] = name
	return path{result}
}
