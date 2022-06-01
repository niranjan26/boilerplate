package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	DataBase DBConfig `json:"database"`
	Host     string   `json:"host"`
	Port     string   `json:"port"`
}

type DBConfig struct {
	Host     string `json:"host"`
	Password string `json:"password"`
	User     string `json:"user"`
	Port     string `json:"port"`
	DBName   string `json:"dbName"`
}

func LoadConfiguration(file string) Config {
	var config Config

	configFile, err := os.Open(file)
	if err != nil {
		log.Fatal("unable to load file", err)
	}

	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)

	err = jsonParser.Decode(&config)
	if err != nil {
		log.Fatal("Unable to parse config", err)
	}

	return config
}
