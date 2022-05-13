package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	DatabaseEnabled   bool
	DatabaseDirectory string
}

func readConfigFile() *os.File {
	homeDirectory, _ := os.UserHomeDir()
	configFile, err := os.Open(homeDirectory + "/.trae")
	if err != nil {
		configFile := createDefaultConfigFile()
		return configFile
	}
	return configFile
}

func createDefaultConfigFile() *os.File {
	homeDirectory, _ := os.UserHomeDir()
	configuration := &Configuration{}
	configuration.DatabaseDirectory = homeDirectory
	configuration.DatabaseEnabled = false
	content, err := json.Marshal(configuration)
	err = os.WriteFile(homeDirectory+"/.trae", content, 0600)
	if err != nil {
		os.Exit(1) // Error
	}
	defaultConfigFile, err := os.Open(homeDirectory + "/.trae")
	return defaultConfigFile
}

func parseConfigFile(configFile *os.File) *Configuration {
	// decoding config file
	decoder := json.NewDecoder(configFile)
	configuration := Configuration{}
	decoderErr := decoder.Decode(&configuration)
	if decoderErr != nil {
		fmt.Println("error:", decoderErr)
	}
	return &configuration
}
