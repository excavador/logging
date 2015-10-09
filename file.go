package logging

import (
	"fmt"
	"os"
)

type File struct {
	path string
	file *os.File
}

type FileArray []File

var flags = os.O_WRONLY | os.O_APPEND | os.O_CREATE

func NewFile(path string) File {
	return File{path, nil}
}

func (self *File) Open() {
	if self.file != nil {
		panic(fmt.Sprintf("log file %v already opened", self.path))
	}
	file, err := os.OpenFile(self.path, flags, 0660)
	if err != nil {
		panic(fmt.Sprintf("log file %v can not be opened %v", self.path, err))
	}
	self.file = file
}

func (self *File) Close() {
	if self.file != nil {
		panic(fmt.Sprintf("file %v already closed", self.path))
	}
	if err := self.file.Close(); err != nil {
		panic(fmt.Sprintf("log file %v can not be closed %v", self.path, err))
	}
	self.file = nil
}

func (self *File) Write(message string) {
	if self.file == nil {
		panic(fmt.Sprintf("log file %v is closed", self.path))
	}
	data := message + "\n"
	self.file.WriteString(data)
}

func NewFileArray(pathes []string) FileArray {
	result := make([]File, len(pathes))
	for index, path := range pathes {
		result[index] = NewFile(path)
	}
	return FileArray(result)
}

func (self *FileArray) Open() {
	for _, file := range []File(*self) {
		file.Open()
	}
}

func (self *FileArray) Close() {
	for _, file := range []File(*self) {
		file.Close()
	}
}

func (self *FileArray) Write(message string) {
	for _, file := range []File(*self) {
		file.Write(message)
	}
}
