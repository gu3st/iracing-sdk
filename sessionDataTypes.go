package irsdk

import (
	"strconv"
	"strings"
)

type SessionVariable struct {
	Value    interface{}
	Unit     string
	Original interface{}
}

func (sv *SessionVariable) UnmarshalYAML(unmarshal func(interface{}) error) error {
	err := unmarshal(&sv.Original)
	if err != nil {
		return err
	}

	sv.Value = sv.Original
	return nil
}

func (t SessionVariable) MarshalYAML() (interface{}, error) {
	return t.Original, nil
}

type Temperature struct {
	Value    float32
	Unit     string
	Original string
}

func (t *Temperature) UnmarshalYAML(unmarshal func(interface{}) error) error {
	err := unmarshal(&t.Original)
	if err != nil {
		return err
	}
	split := strings.Split(t.Original, " ")
	parsedVal, err := strconv.ParseFloat(split[0], 32)

	t.Value = float32(parsedVal)
	t.Unit = split[1]
	return nil
}

func (t Temperature) MarshalYAML() (interface{}, error) {
	return t.Original, nil
}


