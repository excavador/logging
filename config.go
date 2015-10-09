package logging

type LoggerConfig struct {
	// Hierhical name of the logger, like "a.b.c", or "http.request"
	Name string
	// File name for write data.
	File string
	// Level
	Level Level
}

type Config struct {
	// Directory for place log files
	Path string
	// Configuration for particual loggers
	Loggers []LoggerConfig
}
