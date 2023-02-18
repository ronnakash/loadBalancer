package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	config := Parse()
	fmt.Printf("number of servers:%d'\n", len(config.Servers))
	lb := NewLoadBalancer(config)
	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.serveProxy(rw, req)
	}
	portStr := strconv.Itoa(lb.port)
	// http.HandleFunc("/", handleRedirect)
	fmt.Printf("serving requests at localhost:%s\n", portStr)
	go http.ListenAndServe(":"+portStr, http.HandlerFunc(handleRedirect))
	lb.ReadFromCLI()
}
