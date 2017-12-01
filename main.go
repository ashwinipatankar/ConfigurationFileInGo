package main

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Port              int
	Static_Variable   string
	Connection_String string
}

func main() {
	//filename is the path to the json config file
	filename := "config/config.development.json"
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	decoder := json.NewDecoder(file)
	var configuration Configuration
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Println(err)
	}
	configuration.Connection_String = os.Getenv("Connection_String")

	log.Println(configuration.Port)
	log.Println(configuration.Connection_String)
	log.Println(configuration.Static_Variable)
}
