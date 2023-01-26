package main

import (
    "fmt"
    "gopkg.in/yaml.v3"
    "io/ioutil"
)


type Config struct {
    Algorithm 	string
    Servers 	[]ServerParams
    Port        string
}


func parse() Config {
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

    // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    //     var server Server
    //     var index int
    //     if config.Algorithm == "round-robin" {
    //         index = roundRobin(index, len(config.Servers))
    //         server = config.Servers[index]
    //     } else if config.Algorithm == "least-connections" {
    //         server = leastConnections(config.Servers)
    //     }
    //     server.Proxy.ServeHTTP(w, r)
    // })


    // http.ListenAndServe(":8080", nil)
}


