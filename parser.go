package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v3"
)


type Config struct {
    Algorithm 	string
    Servers 	[]ServerParams
    Port        string
}


func ParseYaml() Config {
	var config Config

    // Read config.yaml file
    yamlFile, err := ioutil.ReadFile("config.yaml")
    if err != nil {
        fmt.Println(err.Error())
    }

	// Parse yaml file into config struct
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println(err.Error())
	}
	return config
}


func Parse() Config {
	var config Config

    // Read config.yaml file
    yamlFile, err := ioutil.ReadFile("config.yaml")
    if err != nil {
        fmt.Println(err.Error())
    }

	// Parse yaml file into config struct
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println(err.Error())
	}

    //override config values with command-line args if they exist
    flag.StringVar(&config.Algorithm, "algorithm", config.Algorithm, "Algorithm name")
    flag.StringVar(&config.Port, "port", config.Port, "Server port")
    flag.Parse()

	return config
}
