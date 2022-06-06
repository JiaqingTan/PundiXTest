package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type AppConfig struct {
	Port            int                 `json:"port"`
	AllowedCommands map[string][]string `json:"allowed_commands"`
	SSHHost         string              `json:"ssh_host"`
	SSHPort         int                 `json:"ssh_port"`
	SSHUser         string              `json:"ssh_user"`
	SSHPassword     string              `json:"ssh_password"`
	FXCoredCommand  string              `json:"fxcored_command"`
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