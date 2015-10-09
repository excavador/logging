package logging

import (
	"fmt"
	"os"
)

type file struct {
	path string
	file *os.File
}

type fileArray struct {
	data map[string]file
}

var flags = os.O_WRONLY | os.O_APPEND | os.O_CREATE

func newFile(path string) file {
	return file{path, nil}
}

func (self *file) Open() {
	if self.file != nil {
		panic(fmt.Sprintf("log file %v already opened", self.path))
	}
	file, err := os.OpenFile(self.path, flags, 0660)
	if err != nil {
		panic(fmt.Sprintf("log file %v can not be opened %v", self.path, err))
	}
	self.file = file
}

func (self *file) Close() {
	if self.file != nil {
		panic(fmt.Sprintf("file %v already closed", self.path))
	}
	if err := self.file.Close(); err != nil {
		panic(fmt.Sprintf("log file %v can not be closed %v", self.path, err))
	}
	self.file = nil
}

func (self *file) Reopen() {
	self.Close()
	self.Open()
}

func (self *file) Write(message string) {
	if self.file == nil {
		panic(fmt.Sprintf("log file %v is closed", self.path))
	}
	data := message + "\n"
	self.file.WriteString(data)
}

func (self *fileArray) newFileArray() fileArray {
	data := make(map[string]file)
	return fileArray{data}
}

func (self *fileArray) Configure(pathes []string) {
	var set map[string]bool
	for _, path := range pathes {
		set[path] = true
	}

	for _, file := range self.data {
		path := file.path
		if _, ok := set[path]; ok {
			delete(set, path)
		} else {
			delete(self.data, path)
			file.Close()
		}
	}

	for path, _ := range set {
		file := newFile(path)
		file.Open()
		self.data[path] = file
	}
}

func (self *fileArray) Reopen() {
	for _, file := range self.data {
		file.Reopen()
	}
}

func (self *fileArray) Write(message string) {
	for _, file := range self.data {
		file.Write(message)
	}
}
