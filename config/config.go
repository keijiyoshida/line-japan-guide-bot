package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config represents configuration of an application.
type Config struct {
	HTTP HTTP
}

// New parses the JSON file specified by path,
// creates Config instance and returns it.
func New(path string) (*Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var conf Config

	if err := json.Unmarshal(b, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
