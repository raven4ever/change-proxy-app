package main

import (
	"feynman/files"
	"feynman/utils"
	"flag"
	"log"
)

func main() {
	// Define flags
	configFilePath := flag.String("config", "./config.yml", "path to config file")

	// Parse flags
	flag.Parse()

	// verify if the config file path exists and is a file
	if files.FileExists(*configFilePath) {
		log.Println("Using config file from:", *configFilePath)
	} else {
		log.Fatal("Config file does not exist:", *configFilePath)
	}

	// Load config file from the config.yml path
	config, err := utils.LoadConfig(*configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("URL encoded credentials:", utils.CredentialsToURLString(config))

}
