package logging

import (
	"strings"
)

// logical path of logger, i.e. for logger "http.request" it would be ["http", "request"]
type name struct {
	data []string
}

func (self name) String() string {
	return strings.Join(self.data, ".")
}

func parseName(nameString string) name {
	result := make([]string, 0)
	for _, path := range strings.Split(nameString, ".") {
		if len(path) > 0 {
			result = append(result, path)
		}
	}
	return name{result}
}

func (self *name) GetChild(nameString string) name {
	length := len(self.data)
	result := make([]string, length+1)
	copy(result, self.data)
	result[length] = nameString
	return name{result}
}
