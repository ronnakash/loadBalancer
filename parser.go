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
	fmt.Fprintf(os.Stdout, "%s %d\n", config.Algorithm, config.Port)
	return config
}

func Parse() Config {
	config := ParseYaml()
	//override config values with command-line args if they exist
	flag.StringVar(&config.Algorithm, "algorithm", config.Algorithm, "Algorithm name")
	flag.IntVar(&config.Port, "port", config.Port, "Server Port")
	flag.Parse()
	//check that algorithm and Port are defined
	//and if not, read them from user
	if config.Algorithm == "" {
		fmt.Print("Enter algorithm name: ")
		fmt.Scan(&config.Algorithm)
	}
	if config.Port == 0 {
		fmt.Print("Enter server Port: ")
		fmt.Scan(&config.Port)
	}

	return config
}

func ReadInput(lb LoadBalancer) {
	line := ReadInputLine()
	if len(line) != 2 {
		fmt.Printf("Argument number mismatch")
		return
	}
	command := line[0]
	switch command {
	//TODO: add and remove should take one arg containing
	//		address and Port in format "addr:Port"
	case "add-server":
		lb.AddServer(line[1])
	case "remove-server":
		lb.RemoveServer(line[1])
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
