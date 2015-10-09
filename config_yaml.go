package logging

func (self *Level) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var data string
	unmarshal(&data)
	if result, err := LevelParse(string(data)); err == nil {
		*self = result
		return nil
	} else {
		return err
	}
}

func (self Level) MarshalYAML() (interface{}, error) {
	if self < MAX {
		return self.String(), nil
	} else {
		return nil, ErrorLevelInvalidValue(self)
	}
}

func (self LoggerConfig) MarshalYAML() (interface{}, error) {
	result := make(map[string]string)
	if len(self.File) > 0 {
		result["file"] = self.File
	}
	result["level"] = self.Level.String()
	return result, nil
}
