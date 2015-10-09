package logging

// Every logger has unique hierarchical name, like "a.b.c", or "http.request"
type LoggerConfig struct {
	// File name for write data.
	File string `json:"file,omitempty",yaml:"file"`
	// Level
	Level string `json:"level",yaml:"level"`
}

type Config struct {
	// Directory for place log files
	Directory string `json:"directory",yaml:"directory"`
	// Loggers configuration
	Loggers map[string]LoggerConfig `json:"loggers,omitempty",yaml:"loggers,omitempty"`
}

func NewConfig() Config {
	return Config{"", make(map[string]LoggerConfig)}
}

func (self *Config) InvalidNames() chan string {
	result := make(chan string)
	go func() {
		for nameString, _ := range self.Loggers {
			name := NewName(nameString)
			if name.String() != nameString {
				result <- nameString
			}
		}
		close(result)
	}()
	return result
}
