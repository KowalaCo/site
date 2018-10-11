package api

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Db string `json:"db"`
}

var Config Configuration

func ReadConfig() {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Fatal("Failed to load config.json: ", err)
	}
	Config = configuration
}
