package logging

import (
	"fmt"
	//	"sync/atomic"
)

type Level int
type ErrorLevelInvalidValue int
type ErrorLevelInvalidString string

func (self ErrorLevelInvalidValue) Error() string {
	return fmt.Sprintf("Invalid Level value: %v", self)
}

func (self ErrorLevelInvalidString) Error() string {
	return fmt.Sprintf("Invalid Level string: %v", self)
}

const (
	TRACE Level = iota
	DEBUG
	INFO
	WARN
	ERROR
	CRIT
	MAX
)

var levelString = []string{
	"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "CRIT",
}

func (self Level) String() string {
	if self < MAX {
		return levelString[self]
	} else {
		panic(ErrorLevelInvalidValue(uint8(self)))
	}
}

func LevelParseString(value string) (Level, error) {
	for index, str := range levelString {
		if str == value {
			return Level(index), nil
		}
	}
	return MAX, ErrorLevelInvalidString(value)
}

func (self *Level) Unmarshal(data []byte) error {
	if result, err := LevelParseString(string(data)); err == nil {
		*self = result
		return nil
	} else {
		return err
	}
}

func (self *Level) UnmarshalJSON(data []byte) error {
	return self.Unmarshal(data)
}

func (self *Level) UnmarshalYAML(data []byte) error {
	return self.Unmarshal(data)
}

/*type LevelAtomic atomic.Value

func (self *LevelAtomic) Get() Level {
	return Level((atomic.Value(*self)).Load())
}

func (self *LevelAtomic) Set(level Level) {
	(atomic.Value(*self)).Store(level)
}*/
