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
    var serverParams []ServerParams
	var servers []Server
	var config Config


    // Read config.yaml file
    yamlFile, err := ioutil.ReadFile("config.yaml")
    if err != nil {
        fmt.Println(err.Error())
        return nil
    }


	// Parse yaml file into config struct
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	serverParams = config.Servers

    // Initialize each server
    for i := range serverParams {
		fmt.Printf("server %d: %s:%s\n", i, serverParams[i].Address, serverParams[i].Port )
        servers = append(servers, newSimpleServer(serverParams[i]))
    }

	return servers

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


