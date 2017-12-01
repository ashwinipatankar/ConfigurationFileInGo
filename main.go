package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Port              int
	Static_Variable   string
	Connection_String string
}

func main() {
	fmt.Println("*************************************************")
	fmt.Println("\t\t\tJSON Example\t\t\t")
	fmt.Println("*************************************************")
	jsonExample()
	fmt.Println("*************************************************")
	fmt.Println("\t\t\tGOING Example\t\t\t")
	fmt.Println("*************************************************")
	goingExample()
	fmt.Println("*************************************************")
	fmt.Println("\t\t\tVIPER Example\t\t\t")
	fmt.Println("*************************************************")
	viperExample()
	fmt.Println("*************************************************")
}
func jsonExample() {
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

func goingExample() {
	configuration := Configuration{}
	err := gonfig.GetConf("config/config.production.json", &configuration)
	if err != nil {
		panic(err)
	}

	log.Println(configuration.Port)
	log.Println(configuration.Connection_String)
	log.Println(configuration.Static_Variable)
}

func viperExample() {

}
