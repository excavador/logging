package logging

import (
	"strings"
)

// logical path of logger, i.e. for logger "http.request" it would be ["http", "request"]
type Name struct {
	data []string
}

func (self Name) String() string {
	return strings.Join(self.data, ".")
}

func NewName(nameString string) Name {
	result := make([]string, 0)
	for _, path := range strings.Split(nameString, ".") {
		if len(path) > 0 {
			result = append(result, path)
		}
	}
	return Name{result}
}

func (self *Name) GetChild(nameString string) Name {
	length := len(self.data)
	result := make([]string, length+1)
	copy(result, self.data)
	result[length] = nameString
	return Name{result}
}

func (self *Name) Components() chan string {
	result := make(chan string)
	go func() {
		for _, component := range self.data {
			result <- component
		}
		close(result)
	}()
	return result
}

func (self *Name) Parents() chan Name {
	result := make(chan Name)
	go func() {
		for index, _ := range self.data {
			result <- Name{self.data[:index]}
		}
		close(result)
	}()
	return result
}
