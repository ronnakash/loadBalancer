package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"gopkg.in/yaml.v3"
)

type Config struct {
	algorithm string `yaml:"algorithm"`
	servers   []ServerParams `yaml:"servers"`
	port      int `yaml:"port"`
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
	config := ParseYaml()
	//override config values with command-line args if they exist
	flag.StringVar(&config.algorithm, "algorithm", config.algorithm, "Algorithm name")
	flag.IntVar(&config.port, "port", config.port, "Server port")
	flag.Parse()
	//check that algorithm and port are defined
	//and if not, read them from user
	if config.algorithm == "" {
		fmt.Print("Enter algorithm name: ")
		fmt.Scan(&config.algorithm)
	}
	if config.port == 0 {
		fmt.Print("Enter server port: ")
		fmt.Scan(&config.port)
	}

	return config
}

func ReadInput(lb LoadBalancer) {
	line := ReadInputLine()
	if len(line) == 0 {
		return
	}
	command := line[0]
	switch command {
	//TODO: add and remove should take one arg containing
	//		address and port in format "addr:port"
	case "add-server":
		if len(line) != 3 {
			fmt.Printf("Argument number mismatch")
		} else {
			lb.AddServer(line[1:])
		}
	case "remove-server":
		if len(line) != 3 {
			fmt.Printf("Argument number mismatch")
		} else {
			lb.RemoveServer(line[1:])
		}
	case "algo":
		if len(line) != 2 {
			fmt.Printf("Argument number mismatch")
		} else {
			lb.ChangeAlgorithm(line[1])
		}
	default:
		fmt.Printf("Command %s is invalid\n", command)
	}
}

func ReadInputLine() []string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	// Split the line into an array of words
	return strings.Fields(text)
}
