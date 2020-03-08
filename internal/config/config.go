package config

import (
	"encoding/json"
	"errors"
	"os"
)

//Config holds the settings to setup the paging service
type Config struct {
	Username  string
	Password  string
	SpaceAPI  string
	Callsigns []string
}

func LoadConfig(filepath string) (Config, error) {
	var res Config
	file, err := os.Open(filepath)
	if err != nil {
		return res, errors.New("Error opening file: " + err.Error())
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&res)
	if err != nil {
		return res, errors.New("Error decoding file: " + err.Error())
	}
	return res, nil
}
