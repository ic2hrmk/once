package main

import (
	"io/ioutil"
	"errors"
	"encoding/json"
	"once/once"
)

type Configuration struct {
	Port int
	OnceConfiguration *once.Configuration
}

func NewConfiguration(filename string) (conf *Configuration, err error) {
	confData, err := ioutil.ReadFile(filename)
	if err != nil {
		err = errors.New("Configuration: error reading configuration file ->" + err.Error())
		return
	}

	err = json.Unmarshal(confData, &conf)
	if err != nil {
		err = errors.New("Configuration: error parsing configuration file ->" + err.Error())
		return
	}

	return
}
