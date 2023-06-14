package main

import (
	"feynman/files"
	"feynman/utils"
	"flag"
	"log"
)

func main() {
	// define flags
	configFilePath := flag.String("config", "./config.yml", "path to config file")

	// parse flags
	flag.Parse()

	// verify if the config file path exists and is a file
	if files.FileExists(*configFilePath) {
		log.Println("Using config file from:", *configFilePath)
	} else {
		log.Fatal("Config file does not exist:", *configFilePath)
	}

	// load config file from the config.yml path
	config, err := utils.LoadConfig(*configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// insert the credentials into the proxy URLs
	http_proxy,  err := utils.InsertCredentialsIntoProxyURLs(config)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Using as proxy URL:", http_proxy)

	// edit each file in the config file and add the variables defined in the config file
	for _, file := range config.Files {
		_, err := files.EditFile(file)
		if err != nil {
			log.Fatal(err)
		}
	}
}
