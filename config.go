package main

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

const tomlLocation = "config.toml"

// Configuration is the configuration for the program.
type Configuration struct {
	ShapefileLocation string
}

// GetConfiguration returns the configuration for the program.
func GetConfiguration() (config *Configuration, err error) {
	var b []byte
	if b, err = ioutil.ReadFile(tomlLocation); err == nil {
		_, err = toml.Decode(string(b), &config)
	}
	return
}
