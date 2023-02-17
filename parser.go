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
	Algorithm string
	Servers   []ServerParams
	Port      int
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
	flag.StringVar(&config.Algorithm, "algorithm", config.Algorithm, "Algorithm name")
	flag.IntVar(&config.Port, "port", config.Port, "Server port")
	flag.Parse()
	//check that algorithm and port are defined
	//and if not, read them from user
	if config.Algorithm == "" {
		fmt.Print("Enter algorithm name: ")
		fmt.Scan(&config.Algorithm)
	}
	if config.Port == 0 {
		fmt.Print("Enter server port: ")
		fmt.Scan(&config.Port)
	}

	return config
}

func ReadInput(lb LoadBalancer) {
	line := ReadInputLine()
	//TODO: ensure not empty first
	if len(line) == 0 {
		return
	}
	command := line[0]
	switch command {
	case "add-server":
		lb.AddServer(line[1:])
	case "remove-server":
		lb.RemoveServer(line[1:])
	case "algo":
		lb.ChangeAlgorithm(line[1])
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
