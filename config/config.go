package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type AppConfig struct {
	Port int `json:"port"`
}

func LoadAppConfig(filePath string) (*AppConfig, error) {
	appConfig := new(AppConfig)

	raw, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("Error occurred while reading config file")
		return nil, err
	}

	err = json.Unmarshal(raw, &appConfig)
	if err != nil {
		log.Println("Error occurred while unmarshalling json config file")
		return nil, err
	}

	return appConfig, nil
}