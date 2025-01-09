package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	URL  string `json:"db_url"`
	User string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func GetConfigFilePath() string {
	home, homeErr := os.UserHomeDir()

	if homeErr != nil {
		fmt.Println(homeErr)
		return ""
	}

	filePath := filepath.Join(home, configFileName)
	return filePath
}

func GetConfig() Config {

	var configStruct Config
	return configStruct
}

func ReadConfigFile() Config {

	jsonFilePath := GetConfigFilePath()
	jsonFile, err := os.ReadFile(jsonFilePath)

	if err != nil {
		log.Fatal(err)
	}

	newConfig := GetConfig()
	err = json.Unmarshal(jsonFile, &newConfig)
	if err != nil {
		log.Fatal(err)
	}

	return newConfig
}

func SetName(name string, configStruct *Config) error {
	configStruct.User = name
	data, err := json.Marshal(configStruct)
	if err != nil {
		return err
	}
	filePath := GetConfigFilePath()
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
