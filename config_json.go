package logging

func (self *Level) UnmarshalJSON(data []byte) error {
	data = data[1 : len(data)-1]
	if result, err := LevelParse(string(data)); err == nil {
		*self = result
		return nil
	} else {
		return err
	}
}

func (self Level) MarshalJSON() ([]byte, error) {
	if self < MAX {
		return []byte(`"` + self.String() + `"`), nil
	} else {
		return nil, ErrorLevelInvalidValue(self)
	}
}
